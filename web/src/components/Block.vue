<script setup>
import Button from 'primevue/button';
import Card from 'primevue/card';
import Panel from 'primevue/panel';
import { ref, defineProps } from 'vue';
defineProps(['item']);
</script>
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
    <!-- <template #icons>
        <Button icon="pi pi-cog" severity="secondary" rounded text @click="toggle" />
        <Menu ref="menu" id="config_menu" :model="items" popup />
    </template> -->
    <template #footer>
        <div class="">
          <span class="text-surface-500 dark:text-surface-400">{{item.date}}</span>
          <div class="flex flex-wrap items-center justify-between gap-4">
            <div class="flex items-center gap-2">
                <Button icon="pi pi-user" rounded text></Button>
                <Button icon="pi pi-bookmark" severity="secondary" rounded text></Button>
            </div>
          </div>
        </div>
    </template>
    
    <!-- <p class="m-6">
        Lorem ipsum dolor sit amet, 
    </p> -->
  </Panel>


  <!-- <Button severity="secondary"> -->
    <!-- <Card class="rounded-lg border "> -->
    <!-- <Card>
      <template #title>Simple Card</template>
      <template #content>
          <p class="m-0">
              Lorem ipsum dolor sit amet, co
          </p>
      </template>
  </Card> -->
  <!-- </Button> -->
</template>
