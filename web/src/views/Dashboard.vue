<template>
  <div class="-mx-2">
    <ReminderViewer 
      v-if="reminderVisible"
      :reminder="reminder"
      @close="showReminder(false)"
    ></ReminderViewer>

    <DataView :value="data" :layout="layout">
      <template #list="slotProps">
        <div class="grid grid-nogutter">
          <div v-for="(item, index) in slotProps.items" :key="index" class="col-12">
            <div class="flex flex-column sm:flex-row sm:align-items-center p-4 gap-3" :class="{ 'border-top-1 surface-border': index !== 0 }">
              <div class="md:w-10rem relative">
                <img class="block xl:block mx-auto border-round w-full" :src="`https://primefaces.org/cdn/primevue/images/product/${item.image}`" :alt="item.name" />
                <Tag :value="item.inventoryStatus" :severity="getSeverity(item)" class="absolute" style="left: 4px; top: 4px"></Tag>
              </div>
              <div class="flex flex-column md:flex-row justify-content-between md:align-items-center flex-1 gap-4">
                <div class="flex flex-row md:flex-column justify-content-between align-items-start gap-2">
                  <div>
                    <span class="font-medium text-secondary text-sm">{{ item.category }}</span>
                    <div class="text-lg font-medium text-900 mt-2">{{ item.name }}</div>
                  </div>
                  <div class="surface-100 p-1" style="border-radius: 30px">
                    <div class="surface-0 flex align-items-center gap-2 justify-content-center py-1 px-2" style="border-radius: 30px; box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.04), 0px 1px 2px 0px rgba(0, 0, 0, 0.06)">
                      <span class="text-900 font-medium text-sm">{{ item.rating }}</span>
                      <i class="pi pi-star-fill text-yellow-500"></i>
                    </div>
                  </div>
                </div>
                <div class="flex flex-column md:align-items-end gap-5">
                  <span class="text-xl font-semibold text-900">${{ item.price }}</span>
                  <div class="flex flex-row-reverse md:flex-row gap-2">
                    <Button icon="pi pi-heart" outlined></Button>
                    <Button icon="pi pi-shopping-cart" label="Buy Now" :disabled="item.inventoryStatus === 'OUTOFSTOCK'" class="flex-auto md:flex-initial white-space-nowrap"></Button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>

      <template #grid="slotProps">
        <div class="bg-surface-0 dark:bg-surface-900 bg-gradient-to-tl from-surface-0 to-surface-50 dark:from-black dark:to-surface-700/80">
          <div class=" grid grid-nogutter grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 ">
            <div v-for="(item, i) in slotProps.items" :key="i" class="p-2">
              <Block
                :item="item"
                @click="setReminder(item)"
                class="p-4 min-h-52 flex flex-col"
              />
              <!-- <div class="p-4 border-2 surface-border surface-card rounded-md flex flex-col">
                <div class="relative mx-auto">
                  <img class="rounded-md w-full" :src="`https://primefaces.org/cdn/primevue/images/product/${item.image}`" :alt="item.name" style="max-width: 300px"/>
                  <Tag :value="item.inventoryStatus" :severity="getSeverity(item)" class="absolute" style="left: 4px; top: 4px"></Tag>
                </div>
                <div class="pt-4">
                  <div class="flex flex-row justify-between items-start gap-2">
                      <div>
                          <span class="font-medium text-secondary text-sm">{{ item.category }}</span>
                          <div class="text-lg font-medium text-900 mt-1">{{ item.name }}</div>
                      </div>
                      <div class="surface-100 p-1" style="border-radius: 30px">
                          <div class="surface-0 flex items-center gap-2 justify-center py-1 px-2" style="border-radius: 30px; box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.04), 0px 1px 2px 0px rgba(0, 0, 0, 0.06)">
                              <span class="text-900 font-medium text-sm">{{ item.rating }}</span>
                              <i class="pi pi-star-fill text-yellow-500"></i>
                          </div>
                      </div>
                  </div>
                  <div class="flex flex-col gap-4 mt-4">
                      <span class="text-2xl font-semibold text-900">${{ item.price }}</span>
                      <div class="flex gap-2">
                          <Button icon="pi pi-shopping-cart" label="Buy Now" :disabled="item.inventoryStatus === 'OUTOFSTOCK'" class="flex-auto white-space-nowrap"></Button>
                          <Button icon="pi pi-heart" outlined></Button>
                      </div>
                  </div>
                </div>
              </div> -->
            </div>
            <!-- <div v-for="(item, index) in slotProps.items" :key="index" class="col-12 sm:col-6 md:col-4 xl:col-6 p-2">
                <div class="p-4 border-1 surface-border surface-card border-round flex flex-column">
                    <div class="surface-50 flex justify-content-center border-round p-3">
                        <div class="relative mx-auto">
                            <img class="border-round w-full" :src="`https://primefaces.org/cdn/primevue/images/product/${item.image}`" :alt="item.name" style="max-width: 300px"/>
                            <Tag :value="item.inventoryStatus" :severity="getSeverity(item)" class="absolute" style="left: 4px; top: 4px"></Tag>
                        </div>
                    </div>
                    <div class="pt-4">
                        <div class="flex flex-row justify-content-between align-items-start gap-2">
                            <div>
                                <span class="font-medium text-secondary text-sm">{{ item.category }}</span>
                                <div class="text-lg font-medium text-900 mt-1">{{ item.name }}</div>
                            </div>
                            <div class="surface-100 p-1" style="border-radius: 30px">
                                <div class="surface-0 flex align-items-center gap-2 justify-content-center py-1 px-2" style="border-radius: 30px; box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.04), 0px 1px 2px 0px rgba(0, 0, 0, 0.06)">
                                    <span class="text-900 font-medium text-sm">{{ item.rating }}</span>
                                    <i class="pi pi-star-fill text-yellow-500"></i>
                                </div>
                            </div>
                        </div>
                        <div class="flex flex-column gap-4 mt-4">
                            <span class="text-2xl font-semibold text-900">${{ item.price }}</span>
                            <div class="flex gap-2">
                                <Button icon="pi pi-shopping-cart" label="Buy Now" :disabled="item.inventoryStatus === 'OUTOFSTOCK'" class="flex-auto white-space-nowrap"></Button>
                                <Button icon="pi pi-heart" outlined></Button>
                            </div>
                        </div>
                    </div>
                </div>
            </div> -->
          </div>
        </div>
      </template>
    </DataView>
  </div>
</template>
<script setup>
import DataView from 'primevue/dataview';
import Button from 'primevue/button';
import Tag from 'primevue/tag';
import Block from '../components/Block.vue';
import ReminderViewer from '../components/ReminderViewer.vue';
import axios from 'axios';
import { useQueryClient, useQuery } from '@tanstack/vue-query';
import { inject, ref } from 'vue';

const { layout } = inject('dashboardLayout');
const queryClient = useQueryClient()
const getTodos = async() => {
  try {
    let res =  (await axios.get('/reminders'))
    if (res.status == 200 && Array.isArray(res.data)) {
      return res.data
    } else {
      return []
    }
  } catch (error) {
    console.log(error);
  }
}
const { isLoading, isSuccess, isPending, isError, isFetching, data, error, refetch } = useQuery({
  queryKey: ['reminders'],
  queryFn: getTodos,
  throwOnError: (err) => console.log(err),
}, queryClient);

const reminder = ref({});
const reminderVisible = ref(false);
const showReminder = (show) => reminderVisible.value = show;
const setReminder = (item) => {
  reminder.value = item;
  showReminder(true);
}
const getSeverity = (product) => {
  switch (product.inventoryStatus) {
    case 'INSTOCK':
      return 'success';
    case 'LOWSTOCK':
      return 'warning';
    case 'OUTOFSTOCK':
      return 'danger';
    default:
      return null;
  }
}
</script>