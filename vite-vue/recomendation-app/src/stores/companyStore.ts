import { defineStore } from "pinia";

export const useCompanyStore = defineStore('companies',{
    state:() => ({data :[]}),
    getters:{},
    actions:{
        async getCompanies(){
            const response = await fetch("http://localhost:8081/companies", {mode:'cors'});
            this.data = await response.json()
        },
    }
})