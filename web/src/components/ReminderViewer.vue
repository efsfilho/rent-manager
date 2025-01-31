<template>
  <div class="grid justify-items-center">
    <Dialog 
      modal
      v-model:visible="visible"
      :header="props.reminder.rent_name"
      :style="{ width: '25rem' }"
    >
      <span class="text-surface-500 dark:text-surface-400 block mb-8">
        Due: {{ formatDate(props.reminder.date) }}
      </span>
      <Divider />
      <template #footer>
        <div class="flex w-full justify-between">
          <div v-if="detail.fetchStatus.value === 'fetching'" class="flex gap-4">
            <ProgressSpinner style="height: 30px" strokeWidth="3"/>
          </div>
          <div v-else class="flex gap-3">
            <div v-for="(item, value) in detail.data.value">
              <Button
                aria-label="Filter"
                :label="getMonthName(item.date)"
                :severity="getButtonStatus(item.status)"
                variant="outlined"
              />
            </div>
          </div>
          <div class="flex gap-4">
            <Button
              :disabled="isPaid"
              :label="isPaid ? 'Paid' : 'Pay'"
              severity="success"
              @click="markAsPaid"
            />
          </div>
        </div>
      </template>
    </Dialog>
  </div>
</template>
<script setup>
import Dialog from 'primevue/dialog';
import Divider from 'primevue/divider';
import ProgressSpinner from 'primevue/progressspinner';
import Button from "primevue/button";
import axios from 'axios';
import { useMutation, useQueryClient, useQuery } from '@tanstack/vue-query';
import { ref, watch, computed } from 'vue';

const emit = defineEmits(['close']);
const props = defineProps(['reminder']);
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
const reminderDate = ref(removeTZ(props.reminder.date));

const queryClient = useQueryClient()
const detail = useQuery({
  queryKey: ['reminderDetail'],
  queryFn: async() => {
    try {
      return (await axios.get('/reminder-detail/'+props.reminder.id)).data;
    } catch (err) {
      console.log(err);
    }
  },
  throwOnError: (err) => {
    console.log(err);
  }
}, queryClient);
const payMutation = useMutation({
  mutationFn: (data) => axios.post('/pay/rent/'+reminderId.value)
});
const mutationOptions = {
  onSuccess: () => emit('close'),
  onSettled: () => queryClient.invalidateQueries({ queryKey: ['reminders'] }),
}
const markAsPaid = () => {
  payMutation.mutate({
    status: 2
  }, mutationOptions);
}
const formatDate = (d) => (new Date(`${d}T00:00:00`)).toLocaleDateString('pt-BR')
const getMonthName = (d) => new Date(d).toLocaleString('pt-BR',{month:'short'}).replace('.','')
const getButtonStatus = (s) => ["info", "warn", "danger", "success"][s];
</script>
