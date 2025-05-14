<script setup>
  import { useCompanyStore } from '../stores/companyStore.ts'

  const companyStore = useCompanyStore();
  companyStore.getCompanies();

  companyStore?.data?.forEach((company) => {
    companyStore.getCompanyRecommendations();
  });
</script>

<template>
  <div class="relative overflow-x-auto shadow-md sm:rounded-lg p-10">
    <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400 border border-gray-200 dark:border-gray-700">
      <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
        <tr>
          <th scope="col" class="px-6 py-3">Company</th>
          <th scope="col" class="px-6 py-3">Ticker</th>
          <th scope="col" class="px-6 py-3">Brokerage</th>
          <th scope="col" class="px-6 py-3">Action</th>
          <th scope="col" class="px-6 py-3">Rating from</th>
          <th scope="col" class="px-6 py-3">Rating to</th>
          <th scope="col" class="px-6 py-3">Target from</th>
          <th scope="col" class="px-6 py-3">Target to</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="company in companyStore?.data" :key="company.company" class="odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700 border-gray-200">
          <td scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ company.company }}</td>
          <td class="px-6 py-4">{{ company.ticker }}</td>
          <td class="px-6 py-4">{{ company.brokerage }}</td>
          <td class="px-6 py-4">{{ company.action }}</td>
          <td class="px-6 py-4">{{ company.rating_from }}</td>
          <td class="px-6 py-4">{{ company.rating_to }}</td>
          <td class="px-6 py-4">{{ company.target_from }}</td>
          <td class="px-6 py-4">{{ company.target_to }}</td>
        </tr>
      </tbody>
    </table>

    <p class="mb-4 text-2xl font-bold leading-none tracking-tight text-gray-900 mt-5">📈 Recommendations:</p>
    <ul class="space-y-2 text-left text-gray-500 dark:text-gray-400">
      <div v-for="company in companyStore?.data" :key="company.company">
        <li class="flex items-center space-x-3 rtl:space-x-reverse" v-if="company.score >= 0.75">
          <svg class="shrink-0 w-3.5 h-3.5 text-green-500 dark:text-green-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 12">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5.917 5.724 10.5 15 1.5"></path>
          </svg>
          <span>{{ company.company }} - Score: <strong>{{ company.score }}</strong></span>
        </li>
      </div>
    </ul>
  </div>
</template>
