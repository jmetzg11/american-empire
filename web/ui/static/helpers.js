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

    // Books expandable rows functionality
    const bookRows = document.querySelectorAll('.book-row');
    if (bookRows.length > 0) {
        bookRows.forEach(row => {
            row.addEventListener('click', function() {
                const index = this.getAttribute('data-book-index');
                const eventsRow = document.getElementById('events-' + index);
                const toggleIcon = this.querySelector('.toggle-icon');

                if (eventsRow.classList.contains('hidden')) {
                    eventsRow.classList.remove('hidden');
                    toggleIcon.textContent = '▼';
                } else {
                    eventsRow.classList.add('hidden');
                    toggleIcon.textContent = '▶';
                }
            });
        });
    }
});
