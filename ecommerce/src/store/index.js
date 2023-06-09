import { createStore } from 'vuex';

const requireContext = require.context('./modules', false, /.*\.js$/)

const modules = requireContext.keys()
    .map(file =>
        [file.replace(/(^.\/)|(\.js$)/g, ''), requireContext(file)]
    )
    .reduce((modules, [name, module]) => {

      if (module.namespaced === undefined) {
        module.namespaced = true
      }

      return { ...modules, [name]: module }
    }, {})

const store = createStore({
    modules
})

export default store

// export default new Vuex.import { createStore } from 'vuex';({
//     modules
// })