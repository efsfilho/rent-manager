<template>
    <div class="m-4">
      <Fieldset legend="Teste 1">
        <p class="m-0">-</p>
        <p class="m-0">
          kjfhdskfjdhfkjdshfdskjfhdskjfhsdaklfjhdkslfgjasfahlkdsjfhdskfjsdhfkjshd
        </p>
        <p class="m-0">-</p>
        <p class="m-0">-</p>
        <p class="m-0">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
        </p>
        <p class="m-0">-</p>
        <BlockInput v-model="data" :disabled="false"/>
      </Fieldset>
    </div>
</template>
<script setup>
import Fieldset from 'primevue/fieldset';
import BlockInput from '../components/BlockInput.vue';
import { ref, watch } from 'vue';
const data = ref();
watch(data, (v) => {
  console.log('data >>>>', JSON.stringify(v))
})
</script>