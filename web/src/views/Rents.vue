<template>
  <div class="card m-4 ">
    <DataTable
      :value="data"
      paginator :rows="5"
      :rowsPerPageOptions="[5, 10, 20, 50]"
      tableStyle="min-width: 50%"
      @row-click="(d) => setRent(d.data)"
      selectionMode="single"
      dataKey="id"
      :metaKeySelection="false"
    >
      <Column field="id" header="Id" style="width: 10%" />
      <Column field="name" header="Name" style="width: 40%" />
      <Column :field="columnFormatDate" header="Date" style="width: 20%" />
      <Column field="done" header="done" style="width: 10%" />
      <Column field="done" header="opt" style="width: 10%" />
    </DataTable>
  </div>
</template>
<script setup>
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import axios from 'axios';
import { useQueryClient, useQuery } from '@tanstack/vue-query';
import { inject } from 'vue';

const { setRent } = inject('openEditRent')
const columnFormatDate = v => (new Date(`${v.date}T00:00:00`)).toLocaleDateString('pt-BR')
const queryClient = useQueryClient()
const { isLoading, isSuccess, isPending, isError, isFetching, data, error, refetch } = useQuery({
  queryKey: ['rents'],
  queryFn: async() => {
    try {
      return (await axios.get('/rent')).data;
    } catch (error) {
      console.log(error);
    }
  },
  throwOnError: (err) => console.log(err)
}, queryClient);

</script>
