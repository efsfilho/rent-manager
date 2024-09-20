<template>
  <!-- <div class=" bg-surface-0 dark:bg-surface-900 bg-gradient-to-tl from-surface-0 to-surface-50 dark:from-black dark:to-surface-700/80 p-8 flex items-center justify-center transition-colors duration-200"> -->
    <!-- <div class="bg-surface-0 dark:bg-surface-900 bg-gradient-to-tl from-surface-0 to-surface-50 dark:from-black dark:to-surface-700/80 transition-colors duration-200"> -->
    <!-- <div class="container sm:px-4 min-h-screen mx-auto rounded-md dark:border-slate-600 sm:border-2 "> -->
        <!-- {{ $route.name }} -->
  <div class="flex items-center justify-center">
    <div class="container border-4 w-full sm:w-11/12 md:w-3/4 lg:w-3/5 xl:w-1/2 2xl:w-2/5 min-h-screen">
      <EditBlock v-if="editBlockVisible" :block=block @close="showEditBlock(false)"></EditBlock>
      <div class="mb-2 sm:m-4 ">
        <Menu class="" />
      </div>
      <div class="m-4 ">
        <router-view class=" " />
      </div>
    </div>
  </div>
</template>
<script setup>
import Menu from "./components/Menu.vue";
import EditBlock from "./components/EditBlock.vue";
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
const changeLayout = () => {
  layout.value = layout.value === 'grid' ? 'list' : 'grid';
}
provide('dashboardLayout', {
  layout,
  changeLayout
});

const editBlockVisible = ref(false);
const block = ref({});
const showEditBlock  = (show) => editBlockVisible.value = show;
const setBlock = (item) => {
  // console.log('item', JSON.stringify(item));
  // blockId.value = item.id;
  block.value = item;
  // block.date = item.date;
  // const date = new Date();
  showEditBlock(true);
}
provide('openEditBlock', {
  showEditBlock,
  setBlock
});
</script>



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