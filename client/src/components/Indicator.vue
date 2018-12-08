<template>

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
        name: "Indicator",
        components: {
          'v-chart': ECharts
        }
    }
</script>

<style scoped>

</style>
