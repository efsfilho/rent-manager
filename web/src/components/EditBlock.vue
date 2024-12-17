<template>
  <div class="grid justify-items-center">
    
    <Dialog v-model:visible="visible"  modal header="Edit Profile" :style="{ width: '25rem' }">
      <template #header>
        <div class="inline-flex items-center justify-center gap-2">
          <span v-if="isANewBlock || edit" class="font-bold whitespace-nowrap">{{ isANewBlock ? 'New': 'Edit' }}</span>
        </div>
      </template>

      <div class="flex items-center gap-4 mb-4">
        <div v-if="isPending" class="flex w-full">
          <ProgressSpinner/>
        </div>
        <div v-else class="w-full">
          <Fluid class="">
            <div class="grid grid-cols-1 gap-8 mt-6">
              <div class="col-span-full">
                <FloatLabel>
                  <label for="new-block">Block name</label>
                  <InputText :disabled="!isANewBlock && !edit" id="new-block" v-model="blockName"/>
                </FloatLabel>
              </div>
              <!-- <div v-else class="col-span-full">
                <p class="px-2.5 py-2 text-slate-500">{{ blockName }}</p>
              </div> -->
              
              <FloatLabel>
                <label for="date">Date</label>
                <DatePicker v-model="blockDate" :disabled="!isANewBlock && !edit" dateFormat="dd/mm/yy" inputId="date" />
              </FloatLabel>
            </div>
          </Fluid>

          <!-- <DatePicker id="datepicker-24h" v-model="datetime24h" showTime hourFormat="24" fluid /> -->
        </div>
      </div>
      <div class="flex items-center gap-4">
        <Button aria-label="Filter" label="10"/>
        <Button severity="secondary" aria-label="Bookmark" label="20" />
        <Button icon="pi pi-search" severity="success" aria-label="Search" />
      </div>

      <template #footer>
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
      </template>
    </Dialog>

  </div>
</template>
<script setup>
import InputText from 'primevue/inputtext';
import FloatLabel from 'primevue/floatlabel';
import Button from "primevue/button";
import ProgressSpinner from 'primevue/progressspinner';
import Dialog from 'primevue/dialog';
import Fluid from 'primevue/fluid';
import DatePicker from 'primevue/datepicker';
import axios from 'axios';
import { useMutation, useQueryClient } from '@tanstack/vue-query';
import { ref, defineEmits, watch, computed,  } from 'vue';

const app_address = import.meta.env.VITE_APP_ADDRESS;

const emit = defineEmits(['close']);
const props = defineProps(['block']);
const isANewBlock = ref(!props.block.id);
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

const blockId = ref(props.block.id);
// 0=pending, 1=due, 2=overdue, 3=paid
// "info", "warn", "error", "success"
const isPaid = ref(props.block.status === 3);
const blockName = ref(props.block.name);
const blockDate = ref(removeTZ(props.block.date));

const createMutation = useMutation({ mutationFn: (data) => axios.post(app_address+'/rent', data) });
const updateMutation = useMutation({ mutationFn: (data) => axios.put(app_address+'/rent/'+blockId.value, {data})});
const deleteMutation = useMutation({ mutationFn: () => axios.delete(app_address+'/rent/'+blockId.value) });
const payMutation = useMutation({ mutationFn: (data) => axios.post(app_address+'/pay/cue/'+blockId.value) });
const queryClient = useQueryClient();
const mutationOptions = {
  onSuccess: () => emit('close'),
  onSettled: () => queryClient.invalidateQueries({ queryKey: ['blocks'] }),
}

const save = () => {
  if (!blockId.value) {
    createMutation.mutate({
      done: false,
      name: blockName.value,
      date: blockDate.value
    }, mutationOptions)
  } else {
    updateMutation.mutate({
      id: blockId.value,
      name: blockName.value,
      date: blockDate.value
    }, mutationOptions);
  }
}

const markAsPaid = () => {
  payMutation.mutate({
    status: 2
  }, mutationOptions);
}

const remove = () => {
  deleteMutation.mutate({
    id: blockId.value,
  }, mutationOptions);
}

const isPending = computed(() => {
  return updateMutation.isPending.value || deleteMutation.isPending.value
});

</script>
