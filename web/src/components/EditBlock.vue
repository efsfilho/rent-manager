<script setup>
import FloatLabel from 'primevue/floatlabel';
import InputText from 'primevue/inputtext';
import Button from "primevue/button";
import ProgressSpinner from 'primevue/progressspinner';
import Dialog from 'primevue/dialog';
import Fluid from 'primevue/fluid';
import axios from 'axios';
import { ref, defineEmits, watch, onMounted, onUnmounted, onBeforeMount,  } from 'vue';
import { useMutation, useQueryClient } from '@tanstack/vue-query';

const emit = defineEmits(['close']);
const props = defineProps(['blockId', 'block']);
const visible = ref(true);
watch(visible, (n) => {
  if (!n) emit('close');
});
const blockText = ref(props.block.text);
const edit = ref(false);

const queryClient = useQueryClient();
const { error, isPending, mutate, reset } = useMutation({
  mutationFn: (newTodo) => axios.put('http://localhost:3000/blocks', newTodo),
  onSuccess: () => emit('close'),
  onSettled: () => queryClient.invalidateQueries({ queryKey: ['blocks'] }),
});

const update = () => {
  mutate({
    id: props.block.id,
    date: new Date(),
    text: blockText.value
  });
}

const remove = () => {
  mutate({
    id: props.block.id,
    date: new Date(),
    text: ''
  });
}

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
                  <InputText id="new-block"  v-model="blockText"/>
                  <label for="new-block">Block name</label>
                </FloatLabel>
              </div>
              <div v-else class="col-span-full">
                <!-- <div class="h-10 rounded-md outline outline-1 outline-slate-300"> -->
                  <p class="px-2.5 py-2 text-slate-500">{{ blockText }}</p>
                <!-- </div> -->
              </div>
            </div>
          </Fluid>
          <!-- <FloatLabel v-else class="w-full">
            <InputText class="w-full" id="new-block"  v-model="blockText" />
            <label for="new-block">Block name</label>
          </FloatLabel> -->
        </div>
      </div>

      <template #footer>
        <div class="flex w-full justify-between ">
          <div class="flex gap-4">
            <Button v-if="edit" label="Delete" severity="secondary" @click="remove()"/>
          </div>
          <div class="flex gap-4">
            <Button v-if="edit" label="Cancel" outlined severity="secondary" @click="edit=false" autofocus />
            <Button v-if="edit" label="Save" outlined severity="secondary" @click="update()"/>
            <Button v-if="!edit" label="Edit" severity="secondary" @click="edit=true"/>
          </div>
        </div>
      </template>
    </Dialog>
  </div>
</template>