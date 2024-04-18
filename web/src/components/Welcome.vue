<template>
  <div class="welcome">
    <div class="container">
      <h1 class="title">{{ title }}-{{ version }}</h1>

      <el-row :gutter="20">
        <el-col :span="8">
          <div class="grid-content">
            <div class="item-title">
              <div><i class="iconfont icon-quick-start"></i></div>
              <div>小试牛刀</div>
            </div>

            <div class="list-box">
              <ul>
                <li v-for="item in samples" :key="item"><a @click="send(item)">{{ item }}</a></li>
              </ul>
            </div>
          </div>
        </el-col>
        <el-col :span="8">
          <div class="grid-content">
            <div class="item-title">
              <div><i class="iconfont icon-plugin"></i></div>
              <div>插件增强</div>
            </div>

            <div class="list-box">
              <ul>
                <li v-for="item in plugins" :key="item.value"><a @click="send(item.value)">{{ item.text }}</a></li>
              </ul>
            </div>
          </div>
        </el-col>
        <el-col :span="8">
          <div class="grid-content">
            <div class="item-title">
              <div><i class="iconfont icon-control"></i></div>
              <div>能力扩展</div>
            </div>

            <div class="list-box">
              <ul>
                <li v-for="item in capabilities" :key="item">
                  <span v-if="item.value === ''">{{ item.text }}</span>
                  <a @click="send(item.value)" v-else>{{ item.text }}</a>
                </li>
              </ul>
            </div>
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>
<script setup>

import {onMounted, ref} from "vue";
import {httpGet} from "@/utils/http";
import {ElMessage} from "element-plus";

const title = ref(process.env.VUE_APP_TITLE)
const version = ref(process.env.VUE_APP_VERSION)

const samples = ref([
  "-",
])

const plugins = ref([
  {
    value: "-",
    text: "-"
  }
])

const capabilities = ref([
  {
    text: "轻松扮演翻译专家，程序员，AI 女友，文案高手...",
    value: ""
  },
  {
    text: "国产大语言模型支持，百度文心，科大讯飞，ChatGLM...",
    value: ""
  },
  {
    text: "绘画：示例",
    value: "绘画：马斯克开拖拉机，20世纪，中国农村。3:2"
  }
])

onMounted(() => {
  httpGet("/api/config/get?key=system").then(res => {
    title.value = res.data.title
  }).catch(e => {
    ElMessage.error("获取系统配置失败：" + e.message)
  })
})

const emits = defineEmits(['send']);
const send = (text) => {
  emits('send', text)
}
</script>
<style scoped lang="stylus">
.welcome {
  text-align center
  display flex
  justify-content center
  margin-top 8vh

  .container {
    max-width 768px;
    width 100%

    .title {
      font-size: 2.25rem
      line-height: 2.5rem
      font-weight 600
      margin-bottom: 4rem
    }

    .grid-content {
      .item-title {
        div {
          padding 6px 10px;

          .iconfont {
            font-size 24px;
          }
        }
      }

      .list-box {
        ul {
          padding 10px;

          li {
            font-size 14px;
            padding .75rem
            border-radius 5px;
            background-color: rgba(247, 247, 248, 1);

            line-height 1.5
            color #666666

            a {
              cursor pointer
              display block
              width 100%
            }
            margin-top 10px;
          }
        }
      }
    }
  }
}
</style>