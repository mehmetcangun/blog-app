import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import Buefy from 'buefy'
//import vuex from 'vuex'

import 'buefy/dist/buefy.css'

const base = axios.create({
  baseURL: process.env.VUE_APP_URL
})

Vue.prototype.$axios = base

Vue.config.productionTip = false

Vue.use(Buefy, {
  defaultIconPack: 'fas',
  defaultContainerElement: '#app'
})

/*const store = {
  state: {
    isLoggedIn: false
  },
  getters: {
    isLoggedIn: state => state.isLoggedIn
  },
  mutations: {
    SET_LOGIN: (state, newValue) => {
      state.isLoggedIn = newValue
    },
  },
  actions: {
    setLogin: ({commit, state}, newValue) => {
      commit('SET_LOGIN', newValue)
      return state.isLoggedIn
    }
  }
}

Vue.use(vuex)
*/

new Vue({
  render: h => h(App),
}).$mount('#app')
