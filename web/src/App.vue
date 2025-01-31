<template>
  <div class="flex items-center justify-center">
    <div class="container border-4 w-full sm:w-11/12 md:w-3/4 lg:w-3/5 xl:w-1/2 2xl:w-2/5 min-h-screen">
      <EditRent
        v-if="editRentVisible"
        :rent=rent
        @close="showEditRent(false)"
      />
      <div class="mb-2 sm:m-4 ">
        <Menu />
      </div>
      <div class="m-4">
        <router-view />
      </div>
    </div>
  </div>
</template>
<script setup>
import Menu from "./components/Menu.vue";
import EditRent from "./components/EditRent.vue";
import { ref, provide } from 'vue';

const layout = ref('grid');
const changeLayout = () => {
  layout.value = layout.value === 'grid' ? 'list' : 'grid';
}
provide('dashboardLayout', {
  layout,
  changeLayout
});

const editRentVisible = ref(false);
const rent = ref({});
const showEditRent  = (show) => editRentVisible.value = show;
const setRent = (item) => {
  rent.value = item;
  showEditRent(true);
}
provide('openEditRent', {
  showEditRent,
  setRent
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
</style>