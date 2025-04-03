<template>
  <div class="grid justify-items-center">
    <Dialog v-model:visible="visible" modal :header="isNewRent ? 'New': 'Edit'" :style="{ width: '25rem' }">

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
                  <InputText id="new-rent" :disabled="!isNewRent && !isEditing" v-model="rentName" />
                </FloatLabel>
              </div>
              <FloatLabel>
                <label for="date">Date</label>
                <DatePicker inputId="date" :disabled="!isNewRent && !isEditing"  v-model="rentDate" dateFormat="dd/mm/yy" />
              </FloatLabel>
            </div>
          </Fluid>
        </div>
      </div>
      <template #footer>
        <div class="flex flex-col w-full gap-4">

          <div class="flex w-full justify-between">
            <div class="flex gap-4">
              <Button v-if="isEditing" label="Delete" severity="danger" @click="remove()"/>
              <Button v-if="!isNewRent && !isEditing" label="Log" severity="info" @click="() => logVisible = !logVisible"/>
            </div>
            
            <div class="flex gap-4">
              <!-- <Button v-if="!isPaid" label="Paid" severity="success" @click="markAsPaid"/> TODO --> 
              <Button v-if="isNewRent || isEditing" label="Cancel" severity="secondary" @click="cancel()" autofocus />
              <Button v-if="isNewRent || isEditing" label="Save" severity="secondary" @click="save()"/>
              <Button v-if="!isNewRent && !isEditing" label="Edit" severity="secondary" @click="edit()"/>
            </div>
          </div>
          {{ props.rent.status }}
          <div class="flex w-full justify-between">
            <RentLog v-if="logVisible" :rent-id="rentId"></RentLog>
          </div>
          <!--  TODO 
          <div v-if="!isEditing" class="flex items-center gap-4">
            <Button aria-label="Filter" label="10"/>
            <Button severity="secondary" aria-label="Bookmark" label="20" />
            <Button icon="pi pi-search" severity="success" aria-label="Search" />
          </div>
          -->
        </div>
      </template>
    </Dialog>
  </div>
</template>
<script setup>
import RentLog from './RentLog.vue';
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
const isEditing = ref(false);
const visible = ref(true);
const logVisible = ref(false);
watch(visible, (n) => {
  if (!n) {
    emit('close');
  }
});

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

const rentId = ref(props.rent.id);
// 0=pending, 1=due, 2=overdue, 3=paid
// "info", "warn", "error", "success"
const isNewRent = ref(!props.rent.id);
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
const edit = () => {
  isEditing.value = true;
}
const cancel = () => {
  if (isEditing.value) {
    isEditing.value = false;
  } else {
    emit('close');
  }
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
