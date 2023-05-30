import { defineStore } from 'pinia';
import axios from 'axios';
import { Tenant } from './tenants';

const url = 'http://localhost:1323'

export interface Property {
  id: string,
  active: boolean,
  name: string,
  address: string,
  rent_id: string,
  tenant: Tenant
}

export interface Properties extends Array<Property>{}

export const usePropertiesStore = defineStore('propertiesStore', {
  state: () => ({
    rawProperties: <Properties> []
  }),
  
  getters: {
    properties(state) {
      return state.rawProperties;
    },

    getPropertyById(state) {
      const filter = (id: string) => {
        return state.rawProperties.find((item) => item.id === id);
      }
      return filter;
    }
  },

  actions: {
    async add(content: Property) {
      await axios.post(`${url}/properties`, content);
      await this.updateStore();
    },

    async updateStore() {
      this.rawProperties = [];
      const properties = await axios.get(`${url}/properties`);
      if (properties.data) {
        this.rawProperties = [...properties.data];
      }
    },
    
    async update(content: Property) {
      const getIndex = (item: Property) => {
        return item.id === content.id
      }
      const i = this.rawProperties.findIndex(getIndex);
      const { id } = content;
      await axios.put(`${url}/properties/${id}`, content);
      this.rawProperties[i] = content;
    },

    async delete(content: Property) {
      const getIndex = (item: Property) => {
        return item.id === content.id;
      }

      const i = this.properties.findIndex(getIndex);
      const { id } = content;
      await axios.delete(`${url}/properties/${id}`);
      await this.updateStore();
    }
  }
});