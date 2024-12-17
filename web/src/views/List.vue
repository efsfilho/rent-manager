<script setup>
import Fieldset from 'primevue/fieldset';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
// import Row from 'primevue/row';                   // optional

import { ref, onMounted } from 'vue';
// import { CustomerService } from '@/service/CustomerService';

onMounted(() => {
    // CustomerService.getCustomersMedium().then((data) => (customers.value = data));
    customers.value = 
                [{
          id: 1000,
          name: 'James Butt',
          country: {
              name: 'Algeria',
              code: 'dz'
          },
          company: 'Benton, John B Jr',
          date: '2015-09-13',
          status: 'unqualified',
          verified: true,
          activity: 17,
          representative: {
              name: 'Ioni Bowcher',
              image: 'ionibowcher.png'
          },
          balance: 70663
      }]
});

const customers = ref();
import axios from 'axios';
import { useQueryClient, useQuery } from '@tanstack/vue-query';
import { inject } from 'vue';
// import BlockLog from '../components/BlockLog.vue';

const app_address = import.meta.env.VITE_APP_ADDRESS;

// const getSeverity = (product) => {
//   switch (product.inventoryStatus) {
//     case 'INSTOCK':
//       return 'success';
//     case 'LOWSTOCK':
//       return 'warning';
//     case 'OUTOFSTOCK':
//       return 'danger';
//     default:
//       return null;
//   }
// }

// const { layout } = inject('dashboardLayout');
const { setBlock } = inject('openEditBlock')

const getTodos = async() => {
  try {
    return (await axios.get(app_address+'/rent')).data;
  } catch (error) {
    console.log(error);
  }
}

const queryClient = useQueryClient()
const { isLoading, isSuccess, isPending, isError, isFetching, data, error, refetch } = useQuery({
  queryKey: ['blocks'],
  queryFn: getTodos,
  throwOnError: (err) => {
    console.log('ASFASFGGGGGGGGGGG>>>>', err)
  }
}, queryClient);
const columnFormatDate = d => (new Date(`${d}T00:00:00`)).toLocaleDateString('pt-BR')
// @click="setBlock(item)"
const teste = (d) => console.log(JSON.stringify(d))
// {"id":2,"done":false,"status":2,"date":"2024-12-07","name":"a2a22aa2"}
</script>
<template>
  <div class="card m-4 ">
    <DataTable :value="data" paginator :rows="5" :rowsPerPageOptions="[5, 10, 20, 50]" tableStyle="min-width: 50%" @row-click="(d) => setBlock(d.data)" selectionMode="single" dataKey="id" :metaKeySelection="false">
      <!-- @rowSelect="onRowSelect" @rowUnselect="onRowUnselect" -->
      <Column field="id" header="Id" style="width: 10%"></Column>
      <Column field="name" header="Name" style="width: 40%"></Column>
      <Column :field="(d) => columnFormatDate(d.date)" header="Date" style="width: 20%"></Column>
      <Column field="done" header="done" style="width: 10%"></Column>
      <Column field="done" header="opt" style="width: 10%" >
        <!-- <Button icon="pi pi-heart" outlined></Button> -->
      </Column>
    </DataTable>
  </div>
</template>
