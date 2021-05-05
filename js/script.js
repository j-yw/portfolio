(function () {
  emailjs.init("user_0wkILVfPQZdHvtcmmP18V");
})();

window.onload = function () {
  document
    .getElementById("contact-form")
    .addEventListener("submit", function (event) {
      event.preventDefault();
      // generate a five digit number for the contact_number variable
      this.contact_number.value = (Math.random() * 100000) | 0;
      // these IDs from the previous steps
      emailjs.sendForm("service_uh0wqxv", "contact_form", this).then(
        function () {
          alert("Thank you for contacting me, I'll get back to you soon");
        },
        function (error) {
          alear("Something went wrong, please try again...", error);
        }
      );
    });
};
