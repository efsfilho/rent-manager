<template>
  <div class="grid justify-items-center">
    <Dialog v-model:visible="visible"  modal :style="{ width: '25rem' }">

      <div class="">
      <!-- <div class="flex w-full justify-between "> -->
        <DataTable :value="logs" scrollable scrollHeight="500px" >
            <Column field="date" header="Name" style="width: 20%"></Column>
            <Column field="log" header="Country"></Column>
            <!-- <Column field="representative.name" header="Representative"></Column>
            <Column field="company" header="Company"></Column> -->
        </DataTable>
      </div>

      <!-- <template #footer>
        <div class="flex w-full justify-between ">
          <div class="flex gap-4">
            <Button v-if="!isANewBlock && edit" label="Delete" severity="danger" @click="remove()"/>
            <Button v-if="!isANewBlock && !edit" label="Log" severity="info" @click="remove()"/>
          </div>
          <div class="flex gap-4">
            <Button v-if="!isANewBlock && !edit && !isPaid" label="Paid" severity="success" @click="markAsPaid"/>
            <Button v-if="isANewBlock || edit" label="Cancel" severity="secondary" @click="edit=false" autofocus />
            <Button v-if="isANewBlock || edit" label="Save" severity="secondary" @click="save()"/>
            <Button v-if="!isANewBlock && !edit" label="Edit" severity="secondary" @click="edit=true"/>
          </div>
        </div>
      </template> -->
    </Dialog>

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
import axios from 'axios';
import { useMutation, useQueryClient } from '@tanstack/vue-query';
import { ref, defineEmits, watch, computed, onMounted  } from 'vue';

const app_address = import.meta.env.VITE_APP_ADDRESS;

const emit = defineEmits(['close']);
const props = defineProps(['block']);
const logs = ref([
  {date: '20/05/2024', log: 'sdasdas'},
  {date: '20/05/2024', log: '1f3d24hdf6h54fh'},
  {date: '20/05/2024', log: 'dsfsdglkçals'},
  {date: '20/05/2024', log: 'fsdgsdg'},
])

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
