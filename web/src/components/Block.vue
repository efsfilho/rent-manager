<template>
  <!-- <Button
    class="w-full"
    icon="pi pi-search"
    iconPos="top"
    :label="item.name"
    :severity="item.done ? 'primary':'secondary' "
    :outlined="item.done" 
  /> -->
  <Message :severity="getMessageStatus(item.status)" class="flex">
    <template #container>
      <div class="h-44">
        <div class="flex justify-center truncate h-3/4">
          <p class="mt-5">{{item.name}}</p>
        </div>
        <div class="flex flex w-full justify-between">
          <div class="flex flex-row">
            <div class="flex flex-row-reverse md:flex-row gap-2">
              <!-- <Button :severity="getButtonStatus(item.status)" outlined :label="item.date"/> -->
              <Button :severity="getButtonStatus(item.status)" outlined>{{ formatDate(item.date) }}</Button>
            </div>
          </div>
        </div>
      </div>
    </template>
  </Message>
</template>
<script setup>
import Message from 'primevue/message';
import Button from 'primevue/button';

import { ref, defineProps } from 'vue';
defineProps(['item']);
const formatDate = (d) => (new Date(`${d}T00:00:00`)).toLocaleDateString('pt-BR')
// HintedString<"error" | "secondary" | "info" | "success" | "warn" | "contrast">
const getMessageStatus = (s) => ["info", "warn", "error", "success"][s];
const getButtonStatus = (s) => ["info", "warn", "danger", "success"][s];
//   const status = ["info", "error", "success"];
//   return "error" | "secondary" | "info" | "success" | "warn" | "contrast"
// }
</script>
<style scoped>
.box {
    padding: 2rem;
    border-radius: 10px;
    text-align: center;
}
</style>