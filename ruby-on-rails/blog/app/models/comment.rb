class Comment < ApplicationRecord
  belongs_to :article

  validates :status, inclusion: { in: VALID_STATUSES }

  VALID_STATUSES = %w[public private archived]

  def archived?
    status == 'archived'
  end
end
