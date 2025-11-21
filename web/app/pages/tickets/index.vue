<template>
  <div class="w-screen min-h-screen bg-green-primary relative overflow-x-hidden p-4">
    <PurpleRibbonComponent>METAL // PUNK //</PurpleRibbonComponent>
    <div class="font-inter rounded-2xl w-full h-full bg-white">
      <GenericBannerComponent />
      <main class="flex flex-col gap-4 p-4">
        <LargeHeaderComponent>
          Koop een ticket
        </LargeHeaderComponent>
        <p>
          We hebben je mailadres nodig om je jouw ticket te sturen. Vul hem hieronder in:
        </p>
        <p
          v-if="error"
          class="rounded color-white p-2 bg-red-300 outline-2 outline-red-500"
        >
          {{ error }}
        </p>
        <input
          v-model="email"
          type="email"
          placeholder="Je email adres"
          required="true"
          class="outline-purple outline-2 bg-purple-100 rounded p-2"
          @keyup.enter="createTicket"
        >
        <div class="bg-zinc-100">
          <p class="m-2 text-xs">Entree ticket: <span class="font-black">€5,00</span></p>
          <p class="m-2 text-xs">Transactiekosten <span class="font-black">€0,40</span></p>
          <p class="text-xl bg-green-secondary p-2">Totaal: <span class="font-black">€5,40</span></p>
        </div>
        <CallToActionComponent
          :disabled="disabled"
          @click="createTicket"
        >
          Naar betaling
        </CallToActionComponent>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import type Order from '~/types/order';

const email: Ref<string> = ref('');
const disabled: Ref<boolean> = ref(true);
const error: Ref<string> = ref('');

navigateTo('/');

watch(email, (to) => {
  disabled.value = !to.includes('@') || !to.includes('.');
});

async function createTicket() {
  error.value = '';

  if (disabled.value) {
    error.value = 'Je moet een email adres invullen!';
    return;
  }

  try {
    const response = await useClientFetch<Order>('orders', {
      method: 'POST',
      body: {
        email: email.value,
      },
    });

    await navigateTo(response.checkout, {
      external: true,
    })

  } catch (e: unknown) {
    error.value = 'Er is iets fout gegaan. Probeer het later opnieuw.';
    return;
  }
}
</script>
