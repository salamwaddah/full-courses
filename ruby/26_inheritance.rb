class Chef
  def make_chicken
    puts "The chef makes chicken"
  end

  def make_salad
    puts "The chef makes salad"
  end

  def make_spacial_dish
    puts "The chef makes BBQ Ribs"
  end
end

class ItalianChef < Chef
  def make_spacial_dish
    puts "The chef makes Pizza Margherita"
  end

  def make_pasta
    puts "The chef makes pasta"
  end
end

chef = ItalianChef.new
chef.make_spacial_dish
chef.make_pasta
