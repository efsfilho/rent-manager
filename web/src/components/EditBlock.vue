<template>
  <div class="grid justify-items-center">
    
    <Dialog v-model:visible="visible"  modal header="Edit Profile" :style="{ width: '25rem' }">
      <template #header>
        <div class="inline-flex items-center justify-center gap-2">
          <span class="font-bold whitespace-nowrap">  block.mode</span>
        </div>
      </template>

      <div class="flex items-center gap-4 mb-4">
        <div v-if="isPending" class="flex w-full">
          <ProgressSpinner/>
        </div>
        <div v-else class="w-full">
          <Fluid class="">
            <div class="grid grid-cols-1 gap-8 mt-6">
              <div v-if="edit" class="col-span-full">
                <FloatLabel>
                  <InputText id="new-block"  v-model="blockName"/>
                  <label for="new-block">Block name</label>
                </FloatLabel>
              </div>
              <div v-else class="col-span-full">
                <p class="px-2.5 py-2 text-slate-500">{{ blockName }}</p>
              </div>
              
              <FloatLabel>
                <DatePicker v-model="blockDate" dateFormat="dd/mm/yy" inputId="date" />
                <label for="date">Date</label>
            </FloatLabel>
            </div>
          </Fluid>

          <!-- <DatePicker id="datepicker-24h" v-model="datetime24h" showTime hourFormat="24" fluid /> -->
        </div>
      </div>

      <template #footer>
        <div class="flex w-full justify-between ">
          <div class="flex gap-4">
            <Button v-if="edit" label="Delete" severity="secondary" @click="remove()"/>
          </div>
          <div class="flex gap-4">
            <Button v-if="!edit && !blockDone" label="Close block" severity="success" @click="markAsDone"/>
            <Button v-if="edit" label="Cancel" outlined severity="secondary" @click="edit=false" autofocus />
            <Button v-if="edit" label="Save" outlined severity="secondary" @click="update()"/>
            <Button v-if="!edit" label="Edit" severity="secondary" @click="edit=true"/>
          </div>
        </div>
      </template>
    </Dialog>

  </div>
</template>
<script setup>
import FloatLabel from 'primevue/floatlabel';
import InputText from 'primevue/inputtext';
import Button from "primevue/button";
import ProgressSpinner from 'primevue/progressspinner';
import Dialog from 'primevue/dialog';
import Fluid from 'primevue/fluid';
import DatePicker from 'primevue/datepicker';
import axios from 'axios';
import { ref, defineEmits, watch, onMounted, onUnmounted, onBeforeMount, computed,  } from 'vue';
import { useMutation, useQueryClient } from '@tanstack/vue-query';
const app_address = import.meta.env.VITE_APP_ADDRESS;

const emit = defineEmits(['close']);
const props = defineProps(['blockId', 'block']);
const visible = ref(true);
watch(visible, (n) => {
  if (!n) emit('close');
});
const blockId = ref(props.block.id);
const blockDone = ref(props.block.done);
const blockName = ref(props.block.name);
const removeTZ = (dateString) => {
  let isoDate = new Date(dateString);
  // remove timezone offset
  isoDate.setMinutes(isoDate.getMinutes() + isoDate.getTimezoneOffset())
  return isoDate;
}
const blockDate = ref(removeTZ(props.block.date));

const edit = ref(false);

watch(blockDate, (date) => {
  // console.log()
  // remove timezone offset
  // a = new Date()
  // date.setMinutes(date.getMinutes() - date.getTimezoneOffset())
  let newDate = new Date(date.getFullYear(), date.getMonth(), date.getDay(), 0, 0, 0)
  // console.log('ASD ASD >>> ', newDate)
  // blockDate.value = newDate
});

const queryClient = useQueryClient();

const updateFunc = (v) => axios.put(app_address+'/cue/'+blockId.value, { data: v });
const deleteFunc = (v) => axios.delete(app_address+'/cue/'+blockId.value, v);
const deleteMutation = useMutation({ mutationFn: deleteFunc });
const updateMutation = useMutation({
  mutationFn: updateFunc,
  // onSuccess: () => emit('close'),
  onSettled: () => queryClient.invalidateQueries({ queryKey: ['blocks'] }),
});


const update = () => {
  let d = blockDate.value;
  console.log('item', new Date(d));
  // console.log('  tz', JSON.stringify(removeTZ(d)));
  updateMutation.mutate({
    id: blockId.value,
    name: blockName.value,
    date: blockDate.value
  });
}

const markAsDone = () => {
  updateMutation.mutate({
    done: true
  });
}

const remove = () => {
  deleteMutation.mutate({
    id: blockId.value,
  },{
    onSuccess: () => emit('close'),
    onSettled: () => queryClient.invalidateQueries({ queryKey: ['blocks'] }),
  });
}

const datetime24h = ref();
const date = ref();
const isPending= computed(() => {
  return updateMutation.isPending.value || deleteMutation.isPending.value
})
</script>
