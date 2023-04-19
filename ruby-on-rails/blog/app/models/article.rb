class Article < ApplicationRecord
  has_many :comments

  validates :title, presence: true
  validates :body, presence: true, length: { minimum: 10 }
  validates :status, inclusion: { in: VALID_STATUSES }

  VALID_STATUSES = %w[public private archived]

  def archived?
    status == 'archived'
  end
end
