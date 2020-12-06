def tree?(line, offset)
  if offset >= line.size
    dyn_offset = offset % line.size
    line[dyn_offset] == '#'
  else
    line[offset] == '#'
  end
end

def slope_down(right=3, down=1)
  File
    .readlines('./input.txt')
    .reduce([0, 0, 0]) { |(count, offset, i), line|
      next [count, offset, i + 1] unless i % down == 0

      count += 1 if tree?(line.chomp!, offset)
      [count, offset + right, i + 1]
    }
    .first
end

def all_routes(routes)
  routes.reduce([]) do |trees, (right, down)|
    trees << slope_down(right, down)
  end
end

# p slope_down
p all_routes([[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]]).reduce :*


