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
    rawTenants: <Tenants> []
  }),
  // state,
  getters: {
    getTenants(state) {
      return state.rawTenants
    },
    // getTenantById: (state) => {
    //   return (id: number) =>
    //     state.items.find((item) => item.id === id);
    //     // state.items.find((item) => !!item && (item as Todo).id === id);
    // },
    getTenantById(state) {
      const filter = (id: number) => {
        return state.rawTenants.find((item) => item.id === id);
      }
      return filter;
    }
  },

  actions: {
    async add(content: Tenant) {
      await axios.post(`${url}/tenants`, content);
      await this.get()
    },

    async get() {
      this.rawTenants = []
      const tenants = await axios.get(`${url}/tenants`);
      if (tenants.data) {
        this.rawTenants = [...tenants.data];
      }
    },
    
    async update(content: Tenant) {
      const getIndex = (item: Tenant) => {
        return item.id === content.id
      }
      const i = this.rawTenants.findIndex(getIndex);
      const { id } = content
      await axios.put(`${url}/tenants/${id}`, content)
      this.rawTenants[i] = content
    },
  
    async delete(content: Tenant) {
      const getIndex = (item: Tenant) => {
        return item.id === content.id
      }

      const i = this.rawTenants.findIndex(getIndex);
      const { id } = content
      await axios.delete(`${url}/tenants/${id}`)
      await this.get()
    }
  }

});