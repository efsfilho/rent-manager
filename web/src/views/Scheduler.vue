<template>
  <Panel>
    <div class="grid gap-4 grid-cols-2 w-2/5">
      <Button label="Update" severity="secondary" @click="refetch()" :disabled="isFetching">
        <ProgressSpinner v-if="isFetching" style="width: 21px; height: 21px"/>
      </Button>
      <Button label="Run check" severity="secondary" @click="check()" >
        <ProgressSpinner v-if="isPending" style="width: 21px; height: 21px"/>
      </Button>
    </div>
    <div class="flex flex-row">
      <div class="basis-1/3">start</div>
      <div class="basis-1/3">end</div>
    </div>
    <VirtualScroller :items="data" :itemSize="50" showLoader :delay="200" class="border rounded" style="width: 350px; height: 200px">
      <template v-slot:item="{ item, options }">
        <div :class="['flex flex-row p-2', { ' bg-slate-100': options.odd }]">
          <div class="basis-1/2">{{ dateFormat(item.start) }}</div>
          <div class="basis-1/2">{{ dateFormat(item.end) }}</div>
        </div>
      </template>
    </VirtualScroller>
  </Panel>
</template>
<script setup>
const app_address = import.meta.env.VITE_APP_ADDRESS;
import Panel from 'primevue/panel'
import Button from 'primevue/button';
import VirtualScroller from 'primevue/virtualscroller';
import ProgressSpinner from 'primevue/progressspinner';
import axios from 'axios';
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query';
const { mutate, isPending } = useMutation({
  mutationFn: () => axios.post(app_address+'/checkDues')
});
const { data, refetch, isFetching } = useQuery({
  queryKey: ['stats'],
  queryFn: async () => (await axios.get(app_address+'/scheduler/history')).data
});
const dateFormat = (d) => {
  return !d ? "" : new Date(d).toLocaleString("pt-BR")
}
const queryClient = useQueryClient();
const check = () => {
  mutate({}, {
    onSettled: () => queryClient.invalidateQueries({ queryKey: ['stats'] }),
  })
}
</script>