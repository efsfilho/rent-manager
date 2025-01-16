<script setup>

import CascadeSelect from 'primevue/cascadeselect';
import Dropdown from 'primevue/dropdown';
import DataView from 'primevue/dataview';
import Menubar from 'primevue/menubar'
import Card from 'primevue/card';
// import Floatlabel from "./presets/aura/floatlabel";
import FloatLabel from 'primevue/floatlabel';
import Fieldset from 'primevue/fieldset';
import InputIcon from 'primevue/inputicon';
import Iconfield from 'primevue/iconfield';
import InputText from 'primevue/inputtext';
import InputSwitch from "primevue/inputswitch";
import IconField from "primevue/iconfield";
import Password from "primevue/password";
import Button from "primevue/button";
import Divider from 'primevue/divider';
import Panel from 'primevue/panel';
import SplitButton from 'primevue/splitbutton';
import ProgressSpinner from 'primevue/progressspinner';
import DatePicker from 'primevue/datepicker';
import axios from 'axios';
import {ref} from 'vue';
const blockText = ref(null);
import { useRouter, useRoute } from 'vue-router';

const router = useRouter();

import { useQueryClient, useQuery, useMutation } from '@tanstack/vue-query';

const { isPending, isSuccess, mutate } = useMutation({
  mutationFn: (newTodo) => axios.post('/cue', newTodo),
  onSuccess: () => {
    blockText.value = '';
    router.push({ name:'home' });
    console.log('NewBlock > usemutation')
  },
})

function onButtonClick() {
  mutate({ name: blockText.value })
}
// const saveItems = ref([
//     {
//         label: 'Update',
//         icon: 'pi pi-refresh'
//     },
//     {
//         label: 'Delete',
//         icon: 'pi pi-times'
//     }
// ])

</script>
<template>
  <div class="grid justify-items-center">
    <ProgressSpinner v-if="isPending"/>
    <Panel v-else class="w-3/5 ">
      <FloatLabel class="flex">
        <InputText class="w-full" id="new-block" v-model="blockText" />
        <label for="new-block">Block name</label>
      </FloatLabel>
      <DatePicker id="datepicker-24h" v-model="datetime24h" showTime hourFormat="24" fluid />
      <template #footer>
        <div class="flex flex-wrap items-center justify-end gap-4">
          <Button @click="onButtonClick" >Salvar</Button>
        </div>
      </template>
    </Panel>
  </div>
</template>