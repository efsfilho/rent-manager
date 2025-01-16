// import "./style.css";
import "primeicons/primeicons.css";

import { createApp } from "vue";
import { createPinia } from "pinia";
import router from "./router";
import App from "./App.vue";
import axios from "axios";
import PrimeVue from "primevue/config";
// import Button from "primevue/button";
// import Floatlabel from "./presets/wind/floatlabel";
// import StyleClass from "primevue/styleclass";
// import Wind from "./presets/wind";
// import Lara from "./presets/lara";
// import Aura from "./presets/aura";
import Aura from '@primevue/themes/aura';
import { VueQueryPlugin } from '@tanstack/vue-query'

axios.defaults.baseURL = import.meta.env.VITE_API_URL || import.meta.env.BASE_URL;

const app = createApp(App);
// app.use(createPinia());
app.use(VueQueryPlugin);
app.use(router);
// app.directive("styleclass", StyleClass);
app.use(PrimeVue, {
  // unstyled: true,
  // // pt: Wind
  // pt: Aura
  // // pt: Lara
  theme: {
    preset: Aura,
    options:{
      // darkModeSelector: 'system'
      darkModeSelector: '',
    }
  },
});
app.mount("#app");


// https://github.com/atakantepe/primevue-tailwind-wind.git