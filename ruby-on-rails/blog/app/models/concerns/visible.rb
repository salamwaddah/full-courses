module Visible
  extend ActiveSupport::Concern

  VALID_STATUSES = %w[public private archived]

  included do
    validates :status, inclusion: { in: VALID_STATUSES }
  end

  def archived?
    stats == 'archived'
  end
end