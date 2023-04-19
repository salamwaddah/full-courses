class Comment < ApplicationRecord
  include Visible

  belongs_to :article

  validates :status, inclusion: { in: VALID_STATUSES }
end
