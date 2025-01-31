<template>
  <div class="grid justify-items-center">
    <Dialog v-model:visible="visible" modal header="Edit Profile" :style="{ width: '25rem' }">
      <template #header>
        <div class="inline-flex items-center justify-center gap-2">
          <span v-if="isNewRent || edit" class="font-bold whitespace-nowrap">{{ isNewRent ? 'New': 'Edit' }}</span>
        </div>
      </template>
      <div class="flex items-center gap-4 mb-4">
        <div v-if="isPending" class="flex w-full">
          <ProgressSpinner />
        </div>
        <div v-else class="w-full">
          <Fluid class="">
            <div class="grid grid-cols-1 gap-8 mt-6">
              <div class="col-span-full">
                <FloatLabel>
                  <label for="new-rent">Rent</label>
                  <InputText :disabled="!isNewRent && !edit" id="new-rent" v-model="rentName"/>
                </FloatLabel>
              </div>
              <FloatLabel>
                <label for="date">Date</label>
                <DatePicker v-model="rentDate" :disabled="!isNewRent && !edit" dateFormat="dd/mm/yy" inputId="date" />
              </FloatLabel>
            </div>
          </Fluid>
        </div>
      </div>
      <div class="flex items-center gap-4">
        <Button aria-label="Filter" label="10"/>
        <Button severity="secondary" aria-label="Bookmark" label="20" />
        <Button icon="pi pi-search" severity="success" aria-label="Search" />
      </div>
      <ReminderLog v-if="logVisible" :rent-id="rentId"></ReminderLog>
      <template #footer>
        <div class="flex w-full justify-between ">
          <div class="flex gap-4">
            <Button v-if="!isNewRent && edit" label="Delete" severity="danger" @click="remove()"/>
            <Button v-if="!isNewRent && !edit" label="Log" severity="info" @click="() => logVisible = !logVisible"/>
          </div>
          <div class="flex gap-4">
            <Button v-if="!isNewRent && !edit && !isPaid" label="Paid" severity="success" @click="markAsPaid"/>
            <Button v-if="isNewRent || edit" label="Cancel" severity="secondary" @click="edit=false" autofocus />
            <Button v-if="isNewRent || edit" label="Save" severity="secondary" @click="save()"/>
            <Button v-if="!isNewRent && !edit" label="Edit" severity="secondary" @click="edit=true"/>
          </div>
        </div>
      </template>
    </Dialog>
  </div>
</template>
<script setup>
import ReminderLog from './ReminderLog.vue';
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

const emit = defineEmits(['close']);
const props = defineProps(['rent']);
const isNewRent = ref(!props.rent.id);
const edit = ref(false);
const visible = ref(true);
const logVisible = ref(false);
watch(visible, (n) => {
  if (!n) {
    // logVisible.value = false;
    emit('close');
  }
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

const rentId = ref(props.rent.id);
// 0=pending, 1=due, 2=overdue, 3=paid
// "info", "warn", "error", "success"
const isPaid = ref(props.rent.status === 3);
const rentName = ref(props.rent.name);
const rentDate = ref(removeTZ(props.rent.date));

const createMutation = useMutation({ mutationFn: (data) => axios.post('/rent', data) });
const updateMutation = useMutation({ mutationFn: (data) => axios.put('/rent/'+rentId.value, {data})});
const deleteMutation = useMutation({ mutationFn: () => axios.delete('/rent/'+rentId.value) });
const payMutation = useMutation({ mutationFn: (data) => axios.post('/pay/cue/'+rentId.value) });
const queryClient = useQueryClient();
const mutationOptions = {
  onSuccess: () => {
    queryClient.refetchQueries({ queryKey: ['reminders'] })
    emit('close')
  },
  onSettled: () => queryClient.invalidateQueries({ queryKey: ['rents'] }),
}

const save = () => {
  if (!rentId.value) {
    createMutation.mutate({
      done: false,
      name: rentName.value,
      date: rentDate.value
    }, mutationOptions)
  } else {
    updateMutation.mutate({
      id: rentId.value,
      name: rentName.value,
      date: rentDate.value
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
    id: rentId.value,
  }, mutationOptions);
}

const isPending = computed(() => {
  return updateMutation.isPending.value || deleteMutation.isPending.value
});

</script>
