<template>
  <div class="grid justify-items-center">
    <Dialog v-model:visible="visible"  modal :header="props.reminder.rent_name" :style="{ width: '25rem' }">
      <!-- <template #header class="flex w-full justify-between">
        <div class="inline-flex items-center justify-center gap-2">
          <span class="font-bold whitespace-nowrap">{{ props.reminder.rent_name }}</span>
        </div>
      </template> -->

      <!-- <div class="flex w-full inline-flex items-center justify-center gap-2">
        <span class="font-bold whitespace-nowrap">{{ props.reminder.rent_name }}</span>
      </div> -->
      <span class="text-surface-500 dark:text-surface-400 block mb-8">Due: {{ formatDate(props.reminder.date)}}</span>
      <Divider />

      <!-- <div class="flex items-center gap-4 mb-4">
        <div v-if="isPending" class="flex w-full">
          <ProgressSpinner/>
        </div>
        <div v-else class="w-full">
          <Fluid class="">
            <div class="grid grid-cols-1 gap-8 mt-6">
              <div class="col-span-full">
                <FloatLabel>
                  <label for="new-block">Block name</label>
                  <InputText :disabled="!isANewBlock && !edit" id="new-block" v-model="reminderName"/>
                </FloatLabel>
              </div>
            <span class="text-surface-500 dark:text-surface-400 block mb-8">Due: {{ formatDate(props.reminder.date)}}</span>
              <FloatLabel>
                <label for="date">Date</label>
                <DatePicker v-model="reminderDate" :disabled="!isANewBlock && !edit" dateFormat="dd/mm/yy" inputId="date" />
                <DatePicker v-model="reminderDate" :disabled="!edit" dateFormat="dd/mm/yy" inputId="date" />
              </FloatLabel>
            </div>
          </Fluid>
        </div>
      </div> -->
      <!-- <Tag style="border: 2px solid var(--border-color); background: transparent; color: var(--text-color)">
                <div class="flex items-center gap-2 px-1">
                    <img alt="Country" src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png" class="flag flag-it" style="width: 18px" />
                    <span class="text-base">Italy</span>
                </div>
            </Tag> -->
      <!-- <div class="flex items-center gap-4">
        <Button aria-label="Filter" label="10"/>
        <Button severity="secondary" aria-label="Bookmark" label="20" />
        <Button icon="pi pi-search" severity="success" aria-label="Search" />
      </div> -->
      <!-- {{detail.fetchStatus}}
      {{detail.isFetched}}
      {{detail.isLoading}}
      {{detail.isPending}}
      {{detail.status}} -->
      <template #footer>
        <div class="flex w-full justify-between ">
          <!-- <div class="flex gap-4"> -->
            <!-- <Button v-if=" edit" label="Delete" severity="danger" @click="remove()"/> -->
            <!-- <Button v-if="!edit" label="Log" severity="info" @click="remove()"/> -->
            <div v-if="detail.fetchStatus.value === 'fetching'" class="flex gap-4">
              <ProgressSpinner style="height: 30px" strokeWidth="3"/>
            </div>
            <div v-else class="flex gap-3">
              <!-- {{ detail.data }} -->
            <div v-for="(item, value) in detail.data.value">
              <!-- {{ getMonthName(item.date) }} -->
              <Button
                aria-label="Filter"
                :label="getMonthName(item.date)"
                :severity="getButtonStatus(item.status)"
                variant="outlined"
              />
            </div>
              <!-- <Button aria-label="Filter" label="10"/>
              <Button severity="secondary" aria-label="Bookmark" label="20" />
              <Button icon="pi pi-search" severity="success" aria-label="Search" /> -->
            </div>
            
          <!-- </div> -->
          <div class="flex gap-4">
            <Button :disabled="isPaid" :label="isPaid ? 'Paid' : 'Pay'" severity="success" @click="markAsPaid"/>
            <Button v-if="edit" label="Cancel" severity="secondary" @click="edit=false" autofocus />
            <Button v-if="edit" label="Save" severity="secondary" @click="save()"/>
            <!-- <Button v-if="!isANewBlock && !edit" label="Edit" severity="secondary" @click="edit=true"/> -->
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
import Card from 'primevue/card';
import Divider from 'primevue/divider';
import Tag from 'primevue/tag';
import axios from 'axios';
import { useMutation, useQueryClient, useQuery } from '@tanstack/vue-query';
import { ref, watch, computed,  } from 'vue';

const app_address = import.meta.env.VITE_APP_ADDRESS;

const emit = defineEmits(['close']);
const props = defineProps(['reminder']);
const edit = ref(false);
const visible = ref(true);
watch(visible, (n) => {
  if (!n) 
    emit('close');
});

const removeTZ = (dateString) => {
  let utcDate = new Date();
  if (dateString) {
    utcDate = new Date(dateString);
  }
  utcDate.setUTCHours(utcDate.getHours(), utcDate.getMinutes(),0 ,0);
  utcDate.setUTCDate(utcDate.getDate());
  return utcDate;
}

const reminderId = ref(props.reminder.id);
// 0=pending, 1=due, 2=overdue, 3=paid
// "info", "warn", "error", "success"
const isPaid = ref(props.reminder.status === 3);
// const reminderName = ref(props.reminder.name);
const reminderDate = ref(removeTZ(props.reminder.date));

const queryClient = useQueryClient()
// const { isLoading, isSuccess, isError, isFetching, data, error, refetch } = useQuery({
const detail = useQuery({
  queryKey: ['reminderDetail'],
  queryFn: async() => {
    try {
      return (await axios.get(app_address+'/reminder-detail/'+props.reminder.id)).data;
    } catch (error) {
      console.log(error);
    }
  },
  throwOnError: (err) => {
    console.log(error);
  }
}, queryClient);

// const createMutation = useMutation({ mutationFn: (data) => axios.post(app_address+'/rent', data) });
// const updateMutation = useMutation({ mutationFn: (data) => axios.put(app_address+'/rent/'+reminderId.value, {data})});
// const deleteMutation = useMutation({ mutationFn: () => axios.delete(app_address+'/rent/'+reminderId.value) });
const payMutation = useMutation({ mutationFn: (data) => axios.post(app_address+'/pay/rent/'+reminderId.value) });
// const queryClient = useQueryClient();
const mutationOptions = {
  onSuccess: () => emit('close'),
  onSettled: () => queryClient.invalidateQueries({ queryKey: ['reminders'] }),
}

const save = () => {
  if (!reminderId.value) {
    createMutation.mutate({
      done: false,
      name: reminderName.value,
      date: reminderDate.value
    }, mutationOptions)
  } else {
    updateMutation.mutate({
      id: reminderId.value,
      name: reminderName.value,
      date: reminderDate.value
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
    id: reminderId.value,
  }, mutationOptions);
}

const isPending = computed(() => {
  return updateMutation.isPending.value || deleteMutation.isPending.value
});

// console.log('d _ '+JSON.stringify(props.reminder))
const formatDate = (d) => (new Date(`${d}T00:00:00`)).toLocaleDateString('pt-BR')
const getMonthName = (d) => new Date(d).toLocaleString('pt-BR',{month:'short'}).replace('.','')
const getButtonStatus = (s) => ["info", "warn", "danger", "success"][s];
</script>
