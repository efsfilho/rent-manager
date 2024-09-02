<script setup>
import Button from 'primevue/button';
import { ref, defineProps } from 'vue';
defineProps(['item']);
</script>
<template>
  <Button
    class="w-full"
    icon="pi pi-search"
    iconPos="top"
    :label="item.name"
    :severity="item.done ? 'primary':'secondary' "
    :outlined="item.done" 
  />
</template>
