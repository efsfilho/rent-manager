import { defineStore } from 'pinia';
import axios from 'axios';

export interface Property {
  id: number,
  name: string,
  address: string
}

export interface Properties extends Array<Property>{}
let idcount = 2;
const url = 'http://10.0.0.11:1323'
export const usePropertiesStore = defineStore('propertiesStore', {
  state: () => ({
    // rawProperties: <Properties> []
    rawProperties: <Properties> [
      {
        id: 0,
        name: "Casa 1",
        address: "Rua 1"
      },{
        id: 1,
        name: "Casa 2",
        address: "Rua 2"
      },{
        id: 2,
        name: "Casa 3",
        address: "Rua 3"
      }
    ]
  }),
  
  getters: {
    properties(state) {
      return state.rawProperties
    },

    getPropertyById(state) {
      const filter = (id: number) => {
        return state.rawProperties.find((item) => item.id === id);
      }
      return filter;
    }
  },

  actions: {
    async add(content: Property) {
      // await axios.post(`${url}/properties`, content);
      // await this.get();
      idcount++
      content.id = idcount;
      this.$state.rawProperties.push(content);
    },

    async get() {
      // const tenants = await axios.get(`${url}/tenants`);
      // if (tenants.data) {
      //   this.tenants = [...tenants.data];
      // }
    },
    
    async update(content: Property) {
      const getIndex = (item: Property) => {
        return item.id === content.id
      }
      const i = this.rawProperties.findIndex(getIndex);
      const { id } = content
      // await axios.put(`${url}/tenants/${id}`, content)
      this.rawProperties[i] = content
    },

    async delete(content: Property) {
      const getIndex = (item: Property) => {
        return item.id === content.id
      }

      const i = this.properties.findIndex(getIndex);
      const { id } = content
      // await axios.delete(`${url}/properties/${id}`)
      // await this.get()
      this.$state.rawProperties = this.properties.filter(e => e.id !== i)
      // console.log("delete ", i);
      
    }
  }

});