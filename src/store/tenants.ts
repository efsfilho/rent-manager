import { defineStore } from 'pinia';
import axios from 'axios';

export interface Tenant {
  id: number,
  // reg?: string,
  name: string;
  cpf: string;
  rg: string;
  birth_date: number;
}

export interface Tenants extends Array<Tenant>{}
// export type Tenants = Tenant[] | undefined[];
// export interface TenantState {
//   items: Tenants;
// }
// const state = (): TenantState => ({
//   items: [],
// });
const url = 'http://10.0.0.11:1323'
export const useTenantStore = defineStore('tenantStore', {
  state: () => ({
    tenants: <Tenants> []
  }),
  // state,
  getters: {
    getTenants(state) {
      return state.tenants
    },
    // getTenantById: (state) => {
    //   return (id: number) =>
    //     state.items.find((item) => item.id === id);
    //     // state.items.find((item) => !!item && (item as Todo).id === id);
    // },
    getTenantById(state) {
      const filter = (id: number) => {
        return state.tenants.find((item) => item.id === id);
      }
      return filter;
    }
  },

  actions: {
    async get() {
      // await new Promise((resolve) => setTimeout(() => resolve(1), 5000))
      this.tenants = []
      const tenants = await axios.get(`${url}/tenants`);
      // console.log(tenants.data);
      if (tenants.data) {
        this.tenants = [...tenants.data];
      }
    },
    async add(tenant: Tenant) {
      // TODO send to api and wait for it's response
      await axios.post(`${url}/tenants`, tenant);
      // this.tenants.push(tenant);
      await this.get()
    },
    async update(tenant: Tenant) {
      const getIndex = (item: Tenant) => {
        return item.id === tenant.id
      }
      const i = this.tenants.findIndex(getIndex);
      const { id } = tenant
      await axios.put(`${url}/tenants/${id}`, tenant)
      this.tenants[i] = tenant
    },
    async delete(tenant: Tenant) {
      const getIndex = (item: Tenant) => {
        return item.id === tenant.id
      }

      const i = this.tenants.findIndex(getIndex);
      const { id } = tenant
      await axios.delete(`${url}/tenants/${id}`)
      await this.get()
    }
  }

});