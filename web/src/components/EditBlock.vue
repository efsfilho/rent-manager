<script setup>
import FloatLabel from 'primevue/floatlabel';
import InputText from 'primevue/inputtext';
import Button from "primevue/button";
import ProgressSpinner from 'primevue/progressspinner';
import Dialog from 'primevue/dialog';
import Fluid from 'primevue/fluid';
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
const edit = ref(false);

const queryClient = useQueryClient();

const updateFunc = (v) => axios.put(app_address+'/cue/'+blockId.value, { data: v });
const deleteFunc = (v) => axios.delete(app_address+'/cue/'+blockId.value, v);
const updateMutation = useMutation({
  mutationFn: updateFunc,
  onSuccess: () => emit('close'),
  onSettled: () => queryClient.invalidateQueries({ queryKey: ['blocks'] }),
});

const deleteMutation = useMutation({ mutationFn: deleteFunc });
const update = () => {
  updateMutation.mutate({
    id: blockId.value,
    name: blockName.value,
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

const isPending= computed(() => {
  return updateMutation.isPending.value || deleteMutation.isPending.value
})
</script>
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
          <Fluid class="" >
            <div class="grid grid-cols-1 gap-4">
              <div v-if="edit" class="col-span-full">
                <FloatLabel>
                  <InputText id="new-block"  v-model="blockName"/>
                  <label for="new-block">Block name</label>
                </FloatLabel>
              </div>
              <div v-else class="col-span-full">
                <p class="px-2.5 py-2 text-slate-500">{{ blockName }}</p>
              </div>
            </div>
          </Fluid>
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