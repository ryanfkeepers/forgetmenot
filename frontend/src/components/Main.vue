<!-- ============================== Script ============================== -->
<script setup>
import { ref } from 'vue';
import { Data } from '$/go/App';

ServiceCategories()
  .then(result => {
    state.value.services = result;
  })
  .catch(err => {
    events.emit("error", err);
  });

</script>

<!-- ============================== Template ============================== -->
<template>
  <div class="main">
    <div class="lhs">
      <div class="lhs_header">Forget Me Not...</div>
      <div class="lsh_body">
	<div class="lhs_group">Group1</div>
	<div class="lhs_group">Group2</div>
      </div>
    </div>
    <div class="rhs">
      <div class="rhs_group">
	      <ul>
		      <li>A</li>
		      <li>B</li>
		      <li>C</li>
	      </ul>
      </div>
    </div>
  </div>
</template>

<!-- ============================== Style ============================== -->
<style scoped>
.main {
  height: 100%;
  width: 100%;
  display: flex;
}

.lhs {
  display: flex;
  flex-direction: column;
  margin: 16px; 
  background: black;
  border: 1px solid hotpink;
}

.lhs_header {
  font-size: 2rem;
  border: 1px solid teal;
}

.lhs_body {
  display: flex;
  flex-direction: column;
}

.rhs {
  display: flex;
  flex-direction: column;
}
</style>
