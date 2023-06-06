<template>
    <div class="main">
        <div class="title">
            <h1>Ecommerce Applicaiton Using Vue and GoLang</h1>
        </div>
        <div class="link">
            <router-link to="/orders">Navigate to Orders</router-link>
        </div>
    </div>
</template>

<script>
export default {
    name: "MainComponent"
}
</script>


<style scoped>
.main {
    background-color: #f2f2f2;
    margin: 0;
    padding: 0;
    display: flex;
    width: 100%;
    justify-content: center;
    flex-direction: column;

    align-items: center;
    height: 90vh;
    font-family: Arial, sans-serif;
  }
  .title {
    width: 100%;
  }

  h1 {
    font-size: 90px;
    text-align: center;
    color: #333333;
  }
  
</style>