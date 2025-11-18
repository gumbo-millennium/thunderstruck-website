<template>
  <div class="w-screen min-h-screen bg-green-primary relative overflow-x-hidden p-4">
    <PurpleRibbonComponent>METAL // PUNK //</PurpleRibbonComponent>
    <div class="font-inter rounded-2xl w-full h-full bg-white">
      <GenericBannerComponent />
      <main class="flex flex-col gap-4 p-4">
        <LargeHeaderComponent>
          Je bestelling
        </LargeHeaderComponent>
        <div
          class="flex justify-center items-center"
        >
          <img
            v-if="order === undefined || order?.state === OrderState.PENDING"
            class="animate-pulse aspect-square h-14"
            src="~/assets/images/gumbo.webp"
          >
          <p
            v-if="order?.state === OrderState.CANCELLED"
          >
            Deze bestelling is geannuleerd.
          </p>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import type Order from '~/types/order';
import { OrderState } from '~/types/order';

const route = useRoute();
const order: Ref<Order | undefined> = ref(undefined);
const interval: Ref<number> = ref(-1);

const { data, error } = await useApi<Order>(`orders/${route.params.id}`)
if (error.value !== undefined) {
  console.error(error.value);
  navigateTo('/'); // TODO: Uncomment
}

order.value = data.value
if (order.value?.state === OrderState.PAID) {
  navigateTo(`/tickets/${order.value?.id}`);
}

onMounted(() => {
  interval.value = setInterval(async () => {
    if (order.value?.state === OrderState.PENDING) {
      await fetchOrder();
    }

    if (order.value?.state === OrderState.PAID) {
      clearInterval(interval.value);
      navigateTo(`/tickets/${order.value?.id}`);
    }
  }, 1000);
});

async function fetchOrder() {
  const response = await useClientFetch(`orders/${route.params.id}`);
  console.log(response);
}
</script>
