<template>
  <div>     
    <!-- HOME MENU -->
    <Menubar v-if="route.name === 'home'" :model="items" class="" breakpoint="600px">
      <template #end >
        <Button
          :icon="layout === 'grid' ? 'pi pi-list': 'pi pi-th-large'"
          @click="changeLayout"
        />
      </template>
    </Menubar>
    <!-- HOME MENU -->
  
    <Toolbar v-else style="height: 53px; padding-top: 7px;">
      <template #start >
        <Button icon="pi pi-arrow-left" text @click="() => router.push({ name:'home' })"/>
      </template>
      <template #center>
        {{ $route.meta.menuTitle }}
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
import { ref, defineProps, inject } from "vue";
import { useRouter, useRoute } from 'vue-router';
const { layout, changeLayout } = inject('dashboardLayout');
const { setBlock } = inject('openEditBlock')

const router = useRouter();
const route = useRoute();
const items = ref([
  {
    label: 'New',
    icon: 'pi pi-plus',
    command: () => setBlock({})
  },
  {
    label: 'List',
    icon: 'pi pi-wrench',
    command: () => router.push({ name:'list' })
  },
  {
    label: 'Scheduler',
    icon: 'pi pi-wrench',
    command: () => router.push({ name:'scheduler' })
  },
  {
    label: 'Teste 2',
    icon: 'pi pi-star',
    command: () => router.push({ name:'teste2' })
  },
  // {
  //   label: 'Projects',
  //   icon: 'pi pi-search',
  //   items: [
  //     {
  //       label: 'Components',
  //       icon: 'pi pi-bolt'
  //     },
  //     {
  //       label: 'Blocks',
  //       icon: 'pi pi-server'
  //     },
  //     {
  //       label: 'UI Kit',
  //       icon: 'pi pi-pencil'
  //     },
  //     {
  //       label: 'Templates',
  //       icon: 'pi pi-palette',
  //       items: [
  //         {
  //           label: 'Apollo',
  //           icon: 'pi pi-palette'
  //         },
  //         {
  //           label: 'Ultima',
  //           icon: 'pi pi-palette'
  //         }
  //       ]
  //     }
  //   ]
  // },
  // {
  //   label: 'Contact',
  //   icon: 'pi pi-envelope'
  // }
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

