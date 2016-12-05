/**
 * Base class for objects which are event sources.
 */
class Observable {
  on(event, handler) {
    bean.on(this, event, handler)
  }

  fire(event, ...data) {
    bean.fire(this, event, ...data)
  }
}

/**
 * MetricSet encapsulates the set of selected metrics to be graphed
 * and the collected data for that graph.
 */
class MetricSet extends Observable {
  /** parent is a Metrics instance which this MetricSet is a subset of */
  constructor(parent) {
    super()
    this.parent = parent
    this.metrics = new Set()
    this.data = {}

    this.parent.on('update', () => {
      // Capture datapoints for any selected metrics
      for (let name of this.metrics) {
        this.store_point(name)
      }

      this.fire('update', this)
    })

    this.on('update', () => { this.parent.render() } )
  }

  add(metric) {
    if (!this.has(metric)) {
      this.metrics.add(metric)
      this.store_point(metric)
      this.fire('add', metric)
      this.fire('update')
    }
  }

  remove(metric) {
    if (this.has(metric)) {
      this.metrics.delete(metric)
      delete this.data[metric]
      this.fire('remove', metric)
      this.fire('update')
    }
  }

  clear() {
    this.metrics.clear()
    this.data = {}
    this.fire('clear')
    this.fire('update')
  }

  has(metric) {
    return this.metrics.has(metric)
  }

  store_point(metric) {
    let points = this.data[metric] || []
    points.push({
      timestamp: this.parent.timestamp,
      value: this.parent.metrics[metric]
    })
    this.data[metric] = points
  }

  plotly_data() {
    return Object.keys(this.data).map(name => {
      let raw = this.data[name]
      return {
        x: raw.map(point => point.timestamp),
        y: raw.map(point => point.value),
        type: 'scatter',
        name: name
      }
    })
  }
}

/**
 * Metrics encapsulates the set of all metrics and their most recent
 * values, fetched from the /debug/metrics endpoint.
 */
class Metrics extends Observable {
  constructor() {
    super()
    this.metrics = {}
    this.names = []
    this.timestamp = null
    this.period = 5000
    this.filter = null
  }

  value(metric) {
    return this.metrics[metric]
  }

  update() {
    axios.get('/debug/metrics')
      .then((rsp) => {
        this.metrics = rsp.data.metrics
        this.names = Object.keys(this.metrics)
        this.timestamp = new Date()
      })
      .then(() => {
        this.render()
        this.fire('update', this)
      })
  }

  start() {
    stop()
    this.interval = window.setInterval(() => this.update(), this.period)
  }

  stop() {
    if (this.interval) {
      window.clearInterval(this.interval)
      this.interval = null
    }
  }

  is_running() {
    return Boolean(this.interval)
  }

  setFilter(filter) {
    this.filter = filter
    this.render()
    set_url()
  }

  render() {
    let re = new RegExp(this.filter || '.*')
    let metrics = this.names.map(name => {
        return {
          name: name,
          value: this.value(name),
          selected: Data.selected.has(name)
        }
      }).filter(metric => re.exec(metric.name))

    $('#metrics-count').text(this.names.length)
    $('#filtered-count').text(metrics.length)

    $('#metric-table').html(templates.metric_table_template({
      data: metrics
    }))

    let input = $('#filter')
    if (input.val() != this.filter) {
      input.val(this.filter)
    }
  }
}

/* ----------------------------------------------------------------------
   Plotting
   ---------------------------------------------------------------------- */

function plot_line_chart() {
  let data = Data.selected.plotly_data()
  Plotly.newPlot('chart', data, {}, Data.options.plotly)
  Data.plot = $('#chart')[0]
}

function plot_area_chart() {
  let stackedArea = function(traces) {
    traces[0].fill = 'tozeroy'
	for(var i=1; i<traces.length; i++) {
      traces[i].fill = 'tonexty'
	  for(var j=0; j<(Math.min(traces[i]['y'].length, traces[i-1]['y'].length)); j++) {
		traces[i]['y'][j] += traces[i-1]['y'][j];
	  }
	}
	return traces;
  }
  let data = stackedArea(Data.selected.plotly_data())
  Plotly.newPlot('chart', data, {}, Data.options.plotly)
  Data.plot = $('#chart')[0]
}

function plot_bar_chart() {
  let data = Data.selected.plotly_data()
  for (let series of data) {
    series.type = 'bar'
  }
  Plotly.newPlot('chart', data, {}, Data.options.plotly)
  Data.plot = $('#chart')[0]
}

var plot_fn = plot_line_chart

function plot() {
  plot_fn()
}

function set_url() {
  let u = new URI()
  u.removeQuery('metric')
  u.removeQuery('filter')
  if (Data.selected.metrics.size > 0) {
    u.addQuery('metric', Array.from(Data.selected.metrics.values()))
  }
  if (Data.metrics.filter) {
    u.addQuery('filter', Data.metrics.filter)
  }
  history.pushState({}, '', u.toString())
}

function parse_url() {
  let q = new URI().query(true)
  if (q.metric) {
    if (typeof(q.metric) == 'string') {
      Data.selected.add(q.metric)
    } else {
      for (let metric of q.metric) {
        Data.selected.add(metric)
      }
    }
  }
  if (q.filter) {
    Data.metrics.setFilter(q.filter)
  }
}

function compile_template(selector) {
  return Handlebars.compile($(selector).html())
}

function toggle_metric(name) {
  let set = Data.selected
  if (set.has(name)) {
    set.remove(name)
  } else {
    set.add(name)
  }
}

function init() {
  templates = {
    metric_table_template: compile_template('#metric-table-template')
  }
  Data = {
    options: {
      plotly: {
        showLink: false,
        modeBarButtonsToRemove: ['toImage', 'sendDataToCloud'],
        displaylogo: false
      }
    }
  }
  Data.metrics = new Metrics()
  Data.selected = new MetricSet(Data.metrics)
  Data.metrics.update()

  plot()

  Data.selected.on('update', plot)
  Data.selected.on('update', set_url)
  Data.metrics.start()
}

$(document).ready(() => {
  init()

  $(document).on('click', 'tr.metric-row td.name', (e) => {
    toggle_metric(e.target.textContent)
  })
  $(document).on('input', '#filter', (e) => {
    Data.metrics.setFilter(e.currentTarget.value)
  })
  $(document).on('click', '#btn-clear', (e) => {
    Data.selected.clear()
  })
  $(document).on('click', '#btn-pause', (e) => {
    let icon = $('#btn-pause i')
    if (Data.metrics.is_running()) {
      Data.metrics.stop()
      icon.removeClass('fa-pause')
      icon.addClass('fa-play')
    } else {
      Data.metrics.start()
      icon.removeClass('fa-play')
      icon.addClass('fa-pause')
    }
  })
  $(document).on('click', '#btn-line-chart', (e) => {
    $('.chart-button').removeClass('is-active')
    $('#btn-line-chart').addClass('is-active')
    plot_fn = plot_line_chart
    plot()
  })
  $(document).on('click', '#btn-area-chart', (e) => {
    $('.chart-button').removeClass('is-active')
    $('#btn-area-chart').addClass('is-active')
    plot_fn = plot_area_chart
    plot()
  })
  $(document).on('click', '#btn-bar-chart', (e) => {
    $('.chart-button').removeClass('is-active')
    $('#btn-bar-chart').addClass('is-active')
    plot_fn = plot_bar_chart
    plot()
  })
  $(window).resize(plot)

  parse_url()
})
