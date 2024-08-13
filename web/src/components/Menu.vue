<script setup>
import DarkToggle from './DarkToggle.vue';
import ColorPalette from './ColorPalette.vue'
import Menubar from 'primevue/menubar';
import Menu from 'primevue/menu';
import Toolbar from 'primevue/toolbar';
import Button from 'primevue/button';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import InputText from 'primevue/inputtext';
import SplitButton from 'primevue/splitbutton';
import SelectButton from 'primevue/selectbutton';
import FloatLabel from 'primevue/floatlabel';
import Panel from 'primevue/panel';

import { ref, defineProps, inject, watch } from "vue";
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();
const route = useRoute();

const back = () => router.push({ name:'home' })

const teste = ref('');
const props = defineProps(['title', 'newsletterPopup']);
// watch(props, (newV)=> {
//     console.log('LOGSDDF', newV.title);
// });
import { menu } from '../stores/menu';

const items = ref([

    {
        label: 'New',
        icon: 'pi pi-plus',
        // command: menu.newBlock(),
        command: () => router.push({ name:'new-block'})
        // command: () => router.push({ name:'new-block', props: { title: 123 } })
        // command: () => router.push({ name:'new-block', params: { title: 123 } })
        // command: () => console.log('asdasdsdsad')
    },
    {
        label: 'Teste 1',
        icon: 'pi pi-home',
        command: () => router.push({ name:'teste1' })
        // command: () => menu.teste1()
    },
    {
        label: 'Teste 2',
        icon: 'pi pi-star',
        command: () => router.push({ name:'teste2' })
        // command: () => menu.test2()
    },
    {
        label: 'Projects',
        icon: 'pi pi-search',
        items: [
            {
                label: 'Components',
                icon: 'pi pi-bolt'
            },
            {
                label: 'Blocks',
                icon: 'pi pi-server'
            },
            {
                label: 'UI Kit',
                icon: 'pi pi-pencil'
            },
            {
                label: 'Templates',
                icon: 'pi pi-palette',
                items: [
                    {
                        label: 'Apollo',
                        icon: 'pi pi-palette'
                    },
                    {
                        label: 'Ultima',
                        icon: 'pi pi-palette'
                    }
                ]
            }
        ]
    },
    {
        label: 'Contact',
        icon: 'pi pi-envelope'
    }
]);


const backs= ref([
  { label: 'Router', icon: 'pi pi-palette', class: 'w-96 bg-green-300' },
  // { label: 'Programmatic', icon: 'pi pi-link' },
  // { label: 'External', icon: 'pi pi-home' }
]);

const saveItems = ref([
    {
        label: 'Update',
        icon: 'pi pi-refresh'
    },
    {
        label: 'Delete',
        icon: 'pi pi-times'
    }
])
const options = ref([
  { icon: 'pi pi-bars', value: 'list' },
  { icon: 'pi pi-th-large', value: 'grid' },
]);

// const option = ref('list');
const { layout, changeLayout } = inject('dashboardLayout');

</script>

<template>
    <!-- {{ $route.name }} -->
    <div>
        <!-- <div class="card"> -->
            <!-- <Menubar :model="items" class="p-menubar" /> -->
        <!-- </div> -->
        
      <!-- HOME MENU -->
      <!-- <Menubar v-if="route.name === 'home'" :model="items" class="menu-mobile"> -->
      <Menubar v-if="route.name === 'home'" :model="items" class="" breakpoint="600px">
        <!-- <template #buttonicon>
          <Button class="ml-4" icon="pi pi pi-bars" ></Button>
        </template> -->
        <template #end >
          <Button 
            
            :icon="layout === 'grid' ? 'pi pi-list': 'pi pi-th-large'"
            @click="changeLayout"
          />
        </template>
      </Menubar>
      <!-- HOME MENU -->
    
      <Toolbar v-else style="height: 53px; padding-top: 7px;">
        <!-- <div class="pb-3"> -->
          <template #start >
            <Button icon="pi pi-arrow-left" text @click="back"/>
          </template>

          <template #center>
              {{ $route.meta.menuTitle }}
          </template>
        <!-- </div> -->

        <!-- <template #end>
            <SplitButton label="Save" :model="saveItems"></SplitButton>
            <DarkToggle />
        </template> -->
      </Toolbar>


    <!-- <Toolbar>
        <template #start>
            <Button icon="pi pi-plus" class="mr-2" severity="secondary" />
            <Button icon="pi pi-print" class="mr-2" severity="secondary" />
            <Button icon="pi pi-upload" severity="secondary" />
        </template>

        <template #center>
            <IconField iconPosition="left">
                <InputIcon>
                    <i class="pi pi-search" />
                </InputIcon>
                <InputText placeholder="Search" />
            </IconField>
            
        </template>

        <template #end>
            <SplitButton label="Save" :model="saveItems"></SplitButton>
            <DarkToggle />
        </template>
    </Toolbar> -->
    </div>
</template>

<style scoped >

/* .menu-mobile {
  @apply  max-sm:rounded-none max-sm:border-x-0 max-sm:border-t-0 max-sm:h-16;
} */

/* .menu-default {
    @apply border rounded-md;
} */
</style>

