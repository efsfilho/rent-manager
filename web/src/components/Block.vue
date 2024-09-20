<template>
  <!-- <Button
    class="w-full"
    icon="pi pi-search"
    iconPos="top"
    :label="item.name"
    :severity="item.done ? 'primary':'secondary' "
    :outlined="item.done" 
  /> -->
  <Panel>
    <template #header >
      <div class="truncate">
        <p >{{item.name}}</p>
      </div>
    </template>
    <template #footer>
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
  </Panel>
</template>
<script setup>
import Button from 'primevue/button';
import Card from 'primevue/card';
import Panel from 'primevue/panel';
import { ref, defineProps } from 'vue';
defineProps(['item']);
const formatDate = (d) => (new Date([d,"00:00:00"])).toLocaleDateString('pt-BR')
</script>
