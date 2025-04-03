<template>
  <div v-if="isFetching" class="flex w-full"  >
    <ProgressSpinner />
  </div>
  <div v-else class="flex w-full" style="height: 150px">
    <VirtualScroller v-if="data.length > 0" :items="data" :itemSize="50" :delay="200" class="border rounded" style="width: 100%; height: 150px">
      <template v-slot:item="{ item, options }">
        <div :class="['flex flex-row p-2', { ' bg-slate-100': options.odd }]">
          <div class="basis-2/5">{{ dateFormat(removeTZ(item.date)) }}</div>
          <div class="basis-3/5">{{ item.text }}</div>
        </div>
      </template>
    </VirtualScroller>
    <div v-else class="mt-4">
      No data
    </div>
  </div>
</template>
<script setup>
import ProgressSpinner from 'primevue/progressspinner';
import VirtualScroller from 'primevue/virtualscroller';
import axios from 'axios';
import { useQuery } from '@tanstack/vue-query';

const props = defineProps(['rentId']);
const { data, refetch, isFetching, isPending } = useQuery({
  queryKey: ['logs'],
  queryFn: async () => {
    return (await axios.get('/rent/history/'+props.rentId)).data
  }
});

const dateFormat = (d) => {
  return !d ? "" : new Date(d).toLocaleString("pt-BR")
}
const removeTZ = (dateString) => {
  // year selecting issue: https://github.com/primefaces/primevue/issues/6203
  let utcDate = new Date();
  if (dateString) {
    utcDate = new Date(dateString);
  }
  utcDate.setUTCHours(utcDate.getHours(), utcDate.getMinutes(),0 ,0);
  utcDate.setUTCDate(utcDate.getDate());
  return utcDate;
}

</script>
