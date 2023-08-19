
<template>
  <div>
    <section class="background-radial-gradient mt-15 min-h-screen flex flex-col">
      <div class="grid mt-15 justify-center items-center">
        <div class="overflow-hidden container max-w-md whitespace-normal break-words rounded-lg p-4 mt-10 mb-10 bg-neutral-100 shadow-md shadow-black/5 dark:bg-neutral-600 dark:shadow-black/10">
          <div class="bg-gray-100 min-h-screen py-8">
            <div class="container mx-auto px-4">
              <h1 class="text-3xl font-semibold mb-4 text-center">About </h1>
              <p class="text-gray-700 mb-6 text-center">
                search engine for emails from the enron dataset, made with vue and golang
                  <p class="text-gray-700 mb-6 text-center">
                    Made it with ❤️  for <a class="hover:underline underline hover:text-blue-500" href="https://github.com/HusskyAngel">HusskyAngel</a> 
                  </p>
              </p>
              <p class="text-gray-700 text-center">
                Contact me: <a class="hover:underline hover:text-blue-500 underline " href="mailto:jangelpth@gmail.com">jangelpth@gmail.com</a>
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showMessage: true,
    };
  },
  methods: {
    toggleMessage() {
      this.showMessage = !this.showMessage;
    },
  },
};
</script>
<style>
  .background-radial-gradient {
    background-color: hsl(218, 41%, 15%);
    background-image: radial-gradient(650px circle at 0% 0%,
        hsl(218, 41%, 35%) 15%,
        hsl(218, 41%, 30%) 35%,
        hsl(218, 41%, 20%) 75%,
        hsl(218, 41%, 19%) 80%,
        transparent 100%),
      radial-gradient(1250px circle at 100% 100%,
        hsl(218, 41%, 45%) 15%,
        hsl(218, 41%, 30%) 35%,
        hsl(218, 41%, 20%) 75%,
        hsl(218, 41%, 19%) 80%,
        transparent 100%);
  }
</style>
