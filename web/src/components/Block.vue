<template>
  <!-- <Button
    class="w-full"
    icon="pi pi-search"
    iconPos="top"
    :label="item.name"
    :severity="item.done ? 'primary':'secondary' "
    :outlined="item.done" 
  /> -->

  <Message severity="success" class="break-words">
    <!-- <div class="truncate">
      <p >{{item.name}}</p>
    </div>
    <div class="">
      <span class="text-surface-500 dark:text-surface-400">{{ formatDate(item.date) }}</span>
      <div class="flex flex-wrap items-center justify-between gap-4">
        <div class="flex items-center gap-2">
            <Button icon="pi pi-user" rounded text></Button>
            <Button icon="pi pi-bookmark" severity="secondary" rounded text></Button>
        </div>
      </div>
    </div>
        
    <div class="grid justify-items-center gap-2">
      <div>
          <span class="font-medium text-surface-500 dark:text-surface-400 text-sm">{{ item.category }}</span>
          <div class="text-lg font-medium mt-2">{{ item.name }}</div>
      </div>
    </div>
    <div class="flex flex-col md:items-end gap-8">
      <span class="text-xl font-semibold">{{ item.name }}</span>
      <div class="flex flex-row-reverse md:flex-row gap-2">
        <Button  outlined>{{ formatDate(item.date) }}</Button>
      </div>
    </div> -->
    <!-- <div class="min-w-0 flex-col collapsible-header truncate left-align gap-y-2 grey darken-4 white-text">
      <div class="max-w-md min-w-0">
          <div class="min-w-0 flow-text break-words ">History : {{item.name}}</div> 
      </div>
    </div> -->
    <p class="text-lg truncate break-words">{{ item.name }}</p>

    <div class=" flex flex-wrap items-center justify-between gap-4 text-wrap truncate break-words">
      <div class="text-lg truncate break-words">{{ item.name }}</div>
      <div class="font-semibold text-lg ">${{ item.date }}</div>
      <div class="flex justify-end">
          <span>
              <Button icon="pi pi-heart" severity="secondary" outlined />
              <Button icon="pi pi-shopping-cart" class="ml-2"/>
          </span>
      </div>
    </div>
  </Message>
  






  <!-- <div class="border border-surface-200 dark:border-surface-700 rounded m-2  p-4">
    <div class="mb-4">
        <div class="relative mx-auto">
            <Tag severity="success" class="absolute" style="left:5px; top: 5px"/>
        </div>
    </div>
    <div class="mb-4 font-medium">{{ item.name }}</div>
    <div class="flex justify-between items-center">
        <div class="mt-0 font-semibold text-xl">${{ item.date }}</div>
        <span>
            <Button icon="pi pi-heart" severity="secondary" outlined />
            <Button icon="pi pi-shopping-cart" class="ml-2"/>
        </span>
    </div>
  </div> -->



  <!-- <Panel  >
    <template #header >
      <div class="truncate">
        <p >{{item.name}}</p>
      </div>
    </template> -->

    <!-- <div :severity="item.done ? 'primary':'secondary'" :outlined="item.done"></div> -->
    <!-- <IconField>
      <InputIcon class="pi pi-search" />
      <InputText v-model="value1" placeholder="Search" />
    </IconField> -->

    <!-- <div class="flex flex-row md:flex-col justify-between items-start gap-2">
      <div>
          <span class="font-medium text-surface-500 dark:text-surface-400 text-sm">{{ item.category }}</span>
          <div class="text-lg font-medium mt-2">{{ item.name }}</div>
      </div>
      <div class="bg-surface-100 p-1" style="border-radius: 30px">
        <div class="bg-surface-0 flex items-center gap-2 justify-center py-1 px-2" style="border-radius: 30px; box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.04), 0px 1px 2px 0px rgba(0, 0, 0, 0.06)">
          <span class="text-surface-900 font-medium text-sm">{{ item.rating }}</span>
          <i class="pi pi-star-fill text-yellow-500"></i>
        </div>
      </div>
    </div>
    <div class="flex flex-col md:items-end gap-8">
      <span class="text-xl font-semibold">${{ item.name }}</span>
      <div class="flex flex-row-reverse md:flex-row gap-2">
          <Button icon="pi pi-heart" outlined></Button>
          <Button icon="pi pi-shopping-cart" label="Buy Now" :disabled="item.done === 'OUTOFSTOCK'" class="flex-auto md:flex-initial whitespace-nowrap"></Button>
      </div>
    </div> -->

    

    <!-- <template #footer>
        <div class="">
          <span class="text-surface-500 dark:text-surface-400">{{ formatDate(item.date) }}</span>
          <div class="flex flex-wrap items-center justify-between gap-4">
            <div class="flex items-center gap-2">
                <Button icon="pi pi-user" rounded text></Button>
                <Button icon="pi pi-bookmark" severity="secondary" rounded text></Button>
            </div>
          </div>
        </div>
    </template>
  </Panel> -->
  
<!-- 
    <div class="flex justify-center items-center gap-2">
        <div v-ripple class="box" style="border: 1px solid rgba(75, 175, 80, 0.3); --p-ripple-background: rgba(75, 175, 80, 0.3)">Green</div>
        <div v-ripple class="box" style="border: 1px solid rgba(255, 193, 6, 0.3); --p-ripple-background: rgba(255, 193, 6, 0.3)">Orange</div>
        <div v-ripple class="box" style="border: 1px solid rgba(156, 39, 176, 0.3); --p-ripple-background: rgba(156, 39, 176, 0.3)">Purple</div>
    </div> -->


</template>
<script setup>
import Button from 'primevue/button';
import Card from 'primevue/card';
import Panel from 'primevue/panel';
import Message from 'primevue/message';
import Tag from 'primevue/tag';

import { ref, defineProps, onMounted} from 'vue';
defineProps(['item']);
const formatDate = (d) => (new Date([d,"00:00:00"])).toLocaleDateString('pt-BR')
// onMounted = () => {
//   this.$primevue.config.ripple = true;
// }
</script>
<style scoped>
.box {
    padding: 2rem;
    border-radius: 10px;
    text-align: center;
}
</style>