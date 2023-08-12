function AccessSignUpStudentPage() {
    fetch('/student/signup', )
  .then(response => {
    if (response.ok) {
      window.location.href = response.url;
    } else {
      console.error('Yêu cầu không thành công. Mã trạng thái HTTP:', response.status);
    }
  })
  .catch(error => {
    // Xử lý lỗi ở đây
     console.error('Lỗi khi yêu cầu trang mới:', error);
  });

}