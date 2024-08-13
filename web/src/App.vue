<script setup>
import DarkToggle from "./components/DarkToggle.vue";
import ColorPalette from "./components/ColorPalette.vue";
import CascadeSelect from 'primevue/cascadeselect';
import Dropdown from 'primevue/dropdown';
import Menu from "./components/Menu.vue";
import Divider from "primevue/divider";
import DataView from 'primevue/dataview';
import Menubar from 'primevue/menubar'
import Card from 'primevue/card';
// import Floatlabel from "./presets/aura/floatlabel";
import FloatLabel from 'primevue/floatlabel';
import Fieldset from 'primevue/fieldset';
import InputIcon from 'primevue/inputicon';
import Iconfield from 'primevue/iconfield';
import InputText from 'primevue/inputtext';
import InputSwitch from "primevue/inputswitch";
import IconField from "primevue/iconfield";
import Password from "primevue/password";
import Button from "primevue/button";

import { ref,provide} from 'vue';
import { useRouter, useRoute } from 'vue-router';

// const router = useRouter();
// const route = useRoute();

// let value = ref(null);
// let name = ref('');
// let password = ref('');
// let email = ref('');
// provide('dashboardLayout', 'grid');

const layout = ref('grid');

function changeLayout() {
  layout.value = layout.value === 'grid' ? 'list' : 'grid';
}

provide('dashboardLayout', {
  layout,
  changeLayout
});



</script>

<template>
  <!-- <div class=" bg-surface-0 dark:bg-surface-900 bg-gradient-to-tl from-surface-0 to-surface-50 dark:from-black dark:to-surface-700/80 p-8 flex items-center justify-center transition-colors duration-200"> -->
    <!-- <div class="bg-surface-0 dark:bg-surface-900 bg-gradient-to-tl from-surface-0 to-surface-50 dark:from-black dark:to-surface-700/80 transition-colors duration-200"> -->
    <!-- <div class="container sm:px-4 min-h-screen mx-auto rounded-md dark:border-slate-600 sm:border-2 "> -->
        <!-- {{ $route.name }} -->
    <div class="flex items-center justify-center">
      <div class="container border-4 w-full sm:w-11/12 md:w-3/4 lg:w-3/5 xl:w-1/2 2xl:w-2/5 min-h-screen">
        <!-- sm	640px	@media (min-width: 640px) { ... }
        md	768px	@media (min-width: 768px) { ... }
        lg	1024px	@media (min-width: 1024px) { ... }
        xl	1280px	@media (min-width: 1280px) { ... }
        2xl	1536px	@media (min-width: 1536px) { ... } -->
      <div class="mb-2 sm:m-4 ">
        <Menu class="" />
      </div>
      <div class="m-4 ">
        <router-view class=" " />
      </div>
    </div>
  </div>

  <!-- <div class=" bg-surface-0 dark:bg-surface-900 bg-gradient-to-tl from-surface-0 to-surface-50 dark:from-black dark:to-surface-700/80  p-8 flex items-center justify-center transition-colors duration-200">
  <div class="container min-h-screen mx-auto px-4 rounded-md dark:border-slate-600 border-2 ">
    <div class="card m-4 ">
      <Menu />
    </div>

    <div class="card m-4  ">
      <Fieldset legend="Header">
        <p class="m-0 font-semibold">
          Lorem ipsssum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </p>
      </Fieldset>
    </div>

    <div class="card m-4 ">
      <div class="card m-4 ">
      </div>
        <div class=" px-4 rounded-md dark:border-slate-600 border-2 ">
          <div class="flex flex-col gap-2 mx-auto" style="min-height: 16rem; max-width: 20rem">
            <div class="text-center mt-3 mb-3 text-xl font-semibold">Create your account</div>
            <div class="mb-4">
                <IconField>
                    <InputIcon>
                        <i class="pi pi-user" />
                    </InputIcon>
                    <InputText id="input" v-model="name" type="text" placeholder="Name" />
                </IconField>
            </div>
            <div class="mb-4">
                <IconField>
                    <InputIcon>
                        <i class="pi pi-envelope" />
                    </InputIcon>
                    <InputText id="email" v-model="email" type="email" placeholder="Email" />
                </IconField>
            </div>
            <div class="mb-4">
                <Password v-model="password" toggleMask placeholder="Password" class="w-full" />
            </div>
            <div class="m-4">
              <FloatLabel>
                <InputText id="username" v-model="value" />
                <label for="username">Username</label>
              </FloatLabel>
            </div>
          </div>
          <div class="flex pt-4 justify-end">
            <Button label="Next" icon="pi pi-arrow-right" iconPos="right" />
          </div>
        </div>
      </div>
    </div>
  </div>

  <div
  class=" bg-surface-0 dark:bg-surface-900 bg-gradient-to-tl from-surface-0 to-surface-50 dark:from-black dark:to-surface-700/80  p-8 flex items-center justify-center transition-colors duration-200"
  >
  <Menu />
    <div
      class="p-12 text-center shadow-md bg-gradient-to-tl from-surface-0 to-surface-50 dark:from-surface-900 dark:to-surface-700/60 dark:border dark:border-surface-700 rounded-md"
    >
      <h1
        class="text-4xl font-extralight text-surface-700 dark:text-surface-0/80 mb-4"
      >
        PrimeVue Tailwind Wind Preset
      </h1>
      <span class="text-xl font-light text-surface-700 dark:text-surface-0/80">
        "Showcase for components to be added soon."
      </span>
     
    </div>
  </div> -->
</template>

<style >

@import url('https://rsms.me/inter/inter.css');

:root {
  font-family: Inter, sans-serif;
  font-size: 14px;
  font-feature-settings: 'liga' 1, 'calt' 1; /* fix for Chrome */
}

@supports (font-variation-settings: normal) {
  :root { font-family: InterVariable, sans-serif; }
}

@tailwind base;
@tailwind components;
@tailwind utilities;

/* html{
  font-size: 14px;
} */

/* @import url('https://fonts.googleapis.com/css2?family=Roboto:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900&display=swap');
html, body {
  font-family: 'Roboto', sans-serif;
}
#app {
  font-family: 'Roboto', sans-serif;
} */

/* @import url('https://fonts.googleapis.com/css2?family=Roboto+Mono:ital,wght@0,100..700;1,100..700&display=swap');
html, body {
  font-family: "Roboto Mono", monospace;
}
#app {
  font-family: "Roboto Mono", monospace;
} */

/* @import url('https://fonts.googleapis.com/css2?family=Open+Sans:ital,wght@0,300..800;1,300..800&display=swap');
html, body {
  font-family: "Open Sans", sans-serif;
}
#app {
  font-family: "Open Sans", sans-serif;
} */

/* .teste {
  @apply py-2 px-5 bg-violet-500 text-white font-semibold rounded-full shadow-md hover:bg-violet-700 focus:outline-none focus:ring focus:ring-violet-400 focus:ring-opacity-75;
} */



</style>