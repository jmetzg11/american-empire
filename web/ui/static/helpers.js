// Search Modal functionality
document.addEventListener('DOMContentLoaded', () => {
    const searchBtn = document.getElementById('searchBtn');
    const searchModal = document.getElementById('searchModal');
    const closeModal = document.getElementById('closeModal');

    if (searchBtn && searchModal && closeModal) {
        searchBtn.addEventListener('click', () => {
            searchModal.classList.remove('hidden');
        });

        closeModal.addEventListener('click', () => {
            searchModal.classList.add('hidden');
        });

        // Close modal when clicking outside
        searchModal.addEventListener('click', (e) => {
            if (e.target === searchModal) {
                searchModal.classList.add('hidden');
            }
        });

        // Close modal with Escape key
        document.addEventListener('keydown', (e) => {
            if (e.key === 'Escape' && !searchModal.classList.contains('hidden')) {
                searchModal.classList.add('hidden');
            }
        });
    }
});
