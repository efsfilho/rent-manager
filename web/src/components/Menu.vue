<template>
  <div>
    <Menubar v-if="route.name === 'home'" :model="items" class="" breakpoint="600px">
      <template #end>
        <Button
          :icon="layout === 'grid' ? 'pi pi-list': 'pi pi-th-large'"
          @click="changeLayout"
        />
      </template>
    </Menubar>
  
    <Toolbar v-else style="height: 53px; padding-top: 7px;">
      <template #start>
        <Button icon="pi pi-arrow-left" text @click="() => router.push({ name:'home' })"/>
      </template>
      <template #center>
        {{ $route.meta.title }}
      </template>
      <!-- <template #end>
        <SplitButton label="Save" :model="saveItems"></SplitButton>
        <DarkToggle />
      </template> -->
    </Toolbar>
  </div>
</template>
<script setup>
import Menubar from 'primevue/menubar';
import Toolbar from 'primevue/toolbar';
import Button from 'primevue/button';
import { ref, inject } from "vue";
import { useRouter, useRoute } from 'vue-router';

const { layout, changeLayout } = inject('dashboardLayout');
const { setRent } = inject('openEditRent')

const router = useRouter();
const route = useRoute();
const items = ref([
  {
    label: 'New',
    icon: 'pi pi-plus',
    command: () => setRent({})
  },
  {
    label: 'Rents',
    icon: 'pi pi-file-edit',
    command: () => router.push({ name: 'rents' })
  },
  {
    label: 'Settings',
    icon: 'pi pi-wrench',
    command: () => router.push({ name:'settings' })
  },
]);
</script>
<style scoped >

/* .menu-mobile {
  @apply  max-sm:rounded-none max-sm:border-x-0 max-sm:border-t-0 max-sm:h-16;
} */

/* .menu-default {
    @apply border rounded-md;
} */
</style>

