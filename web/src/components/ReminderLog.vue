<template>
  <div v-if="isFetching" class="flex w-full "  >
    <ProgressSpinner />
  </div>
  <div v-else>
    <VirtualScroller v-if="data.length > 0"  :items="data" :itemSize="50" :delay="200" class="border rounded" style="width: 100%; height: 200px">
      <template v-slot:item="{ item, options }">
        <div :class="['flex flex-row p-2', { ' bg-slate-100': options.odd }]">
          <div class="basis-1/3">{{ dateFormat(removeTZ(item.date)) }}</div>
          <div class="basis-2/3">{{ item.text }}</div>
        </div>
      </template>
    </VirtualScroller>
    <div v-else class="mt-4">
      No data
    </div>
  </div>
</template>
<script setup>
import InputText from 'primevue/inputtext';
import FloatLabel from 'primevue/floatlabel';
import Button from "primevue/button";
import ProgressSpinner from 'primevue/progressspinner';
import Dialog from 'primevue/dialog';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Fluid from 'primevue/fluid';
import DatePicker from 'primevue/datepicker';
import VirtualScroller from 'primevue/virtualscroller';
import axios from 'axios';
import { ref, watch } from 'vue';
// import { useMutation, useQueryClient } from '@tanstack/vue-query';
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query';

const { data, refetch, isFetching, isPending } = useQuery({
  queryKey: ['logs'],
  queryFn: async () => {
    // (await axios.get('/scheduler/history')).data
    const d = (await axios.get('/rent/history/'+props.rentId)).data
    // return ['dddd','dss']
    // console.log('DDDDD', d.length)
    // d.length = 10;
    // if (d.length == 0){
    //   d.push({ text:'no data'})
    // }
    return d
  }
});
const dateFormat = (d) => {
  return !d ? "" : new Date(d).toLocaleString("pt-BR")
}

// const emit = defineEmits(['close']);
const props = defineProps(['rentId']);
// const logs = ref(Array.from([
//   {date: '20/05/2024', log: 'sdasdas'},
//   {date: '20/05/2024', log: '1f3d24hdf6h54fh'},
//   {date: '20/05/2024', log: 'dsfsdglkçals'},
//   {date: '20/05/2024', log: 'fsdgsdg'},
// ]))

// onMounted(() => {
//   logs.value = [
//     {
//       id: 1000,
//       name: 'James Butt',
//       country: {
//           name: 'Algeria',
//           code: 'dz'
//       },
//       company: 'Benton, John B Jr',
//       date: '2015-09-13',
//       status: 'unqualified',
//       verified: true,
//       activity: 17,
//       representative: {
//           name: 'Ioni Bowcher',
//           image: 'ionibowcher.png'
//       },
//       balance: 70663
//     },
//   ]
// })
const edit = ref(false);
const visible = ref(true);
watch(visible, (n) => {
  if (!n) 
    emit('close');
});

const removeTZ = (dateString) => {
  // let isoDate = new Date();
  // if (dateString) {
  //   isoDate = new Date(dateString);
  // }
  // // remove timezone offset
  // isoDate.setMinutes(isoDate.getMinutes() + isoDate.getTimezoneOffset())
  // return isoDate;
  // // year selecting issue: https://github.com/primefaces/primevue/issues/6203
  let utcDate = new Date();
  if (dateString) {
    utcDate = new Date(dateString);
  }
  utcDate.setUTCHours(utcDate.getHours(), utcDate.getMinutes(),0 ,0);
  utcDate.setUTCDate(utcDate.getDate());
  return utcDate;
}

</script>
