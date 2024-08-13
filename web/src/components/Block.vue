<script setup>
import Button from 'primevue/button';
import { ref, defineProps } from 'vue';
defineProps(['item']);

</script>
<template>
  <!-- <div class="w-full"> -->
    <!-- <EditBlock v-if="editBlockVisible" @show="showEditBlock" ></EditBlock> -->
    <Button :label="item.text" class="w-full" icon="pi pi-search" iconPos="top" severity="secondary" />
  <!-- </div> -->

</template>
