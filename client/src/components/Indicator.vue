<template>

  <section>
    <figure style="background-color: #003;">
      <v-chart
        :options="loadDataOption"
        ref="bbb"
        theme="vintage"
        :init-options="initOptions"
      />
    </figure>
  </section>

</template>

<script>
  import ECharts from 'vue-echarts/components/ECharts'

  import 'echarts/lib/chart/bar'
  import 'echarts/lib/chart/line'
  import 'echarts/lib/chart/pie'
  import 'echarts/lib/chart/map'
  import 'echarts/lib/chart/radar'
  import 'echarts/lib/chart/scatter'
  import 'echarts/lib/chart/effectScatter'
  import 'echarts/lib/component/tooltip'
  import 'echarts/lib/component/polar'
  import 'echarts/lib/component/geo'
  import 'echarts/lib/component/legend'
  import 'echarts/lib/component/title'
  import 'echarts/lib/component/visualMap'
  import 'echarts/lib/component/dataset'
  // built-in theme
  import 'echarts/theme/dark'
  import 'echarts/theme/vintage'

  export default {
    name: 'Indicator',
    components: {
      'v-chart': ECharts
    },
    data () {
      var data111 = [[1, 2], [2, 12], [3, 21], [4, 222], [5, 11]]
      var data222 = [[1, 21], [2, 11], [3, 20]]

// var worldMapContainer = document.getElementById('main');

// //用于使chart自适应高度和宽度,通过窗体高宽计算容器高宽
// var resizeWorldMapContainer = function () {
//   worldMapContainer.style.width = window.innerWidth+'px';
//   worldMapContainer.style.height = window.innerHeight+'px';
// };
// //设置容器高宽
// resizeWorldMapContainer();

      return {
        // data:[1,2,3],
        initOptions: {
          width: (window.innerWidth - 10) + 'px',
          height: (window.innerHeight - 20) + 'px'
        },
        websocket: null,
        loadDataOption: null
      }
    },

    created () {
      // 页面刚进入时开启长连接
      // this.initWebSocket();
      // console.log(v-chart.getWidth());
      this.httpGetTrend()
    // this.loadData();
    },
    destroyed: function () {
      // 页面销毁时关闭长连接
      this.websocketclose()
    },
    methods: {
      initWebSocket () { // 初始化weosocket
        const wsuri = 'ws://121.40.165.18:8800'
        this.websock = new WebSocket(wsuri)
        this.websock.onopen = this.websocketonopen

        this.websock.onerror = this.websocketonerror

        this.websock.onmessage = this.websocketonmessage
        this.websock.onclose = this.websocketclose
      },

      websocketonopen () {
        console.log('WebSocket连接成功')
      },
      websocketonerror (e) { // 错误
        console.log('WebSocket连接发生错误')
      },
      websocketonmessage (e) { // 数据接收
        // const redata = JSON.parse(e.data);
        console.log(e.data)
      },

      websocketsend (agentData) { // 数据发送
        this.websock.send(agentData)
      },

      websocketclose (e) { // 关闭
        console.log('connection closed (' + e.code + ')')
      },

      getKlines () {
        this.$http.get('/api/user/info', {headers: auth.getAuthHeader()})
          .then(response => {
            this.userinfo = response.body
          }, response => {
            if (response.status === 401) {
              auth.logout(this)
            }
            console.log(response)
          })
      },
      loadData () {
        this.loadDataOption = ({
          title: {
            text: '走势图',
            // left: 0,
            subtext: '类GBI(内参中货币权重后走势)'
          },
          legend: {
            data: ['走势']
            // height: 800,
            // width: 12000
          },
          dataZoom: [{
            type: 'slider',
            show: true,
            realtime: false,
            start: 80,
            end: 100,
            // filterMode: 'none',
            xAxisIndex: [0]
          },
          {
            type: 'inside',
              // filterMode: 'none',
              // start: 30,
              // end: 70,
            xAxisIndex: [0]
          }],
          toolbox: {
            show: true,
            feature: {
              dataZoom: {
                yAxisIndex: 'none'
              },
              dataView: {readOnly: false},
              magicType: {type: ['line', 'bar']},
              restore: {},
              saveAsImage: {}
            }
          },

          tooltip: {
            trigger: 'axis',
            axisPointer: {
              type: 'cross',
              axis: 'x'
            }
          },
          xAxis: {
            type: 'category',
            data: this.data111.map(function (item) {
              return item[0]
            }),
            boundaryGap: false
          },
          yAxis: [
            {
              type: 'value',
              //  axisLabel: {
              //     formatter: '{value} '
              // },
              min: function (value) {
                return value.min * 0.99
              },
              max: function (value) {
                return value.max * 1.01
              }

            }
            // {
            //  type: 'value',
            //     axisLabel: {
            //        formatter: '{value} %'
            //    }
            //  }
          ],
          series: [
            {
              name: '走势',
              type: 'line',
              // xAxisIndex: 0,
              yAxisIndex: 0,
              data: this.data111,
              smooth: true,
              markPoint: {
                data: [{type: 'max', name: '最大值'}, {type: 'min', name: '最小值'}]
              },
              markLine: {
                data: [{type: 'average', name: '平均值'}]
              }
            }]
        })
      }
    }
  }
</script>

<style scoped>

</style>
