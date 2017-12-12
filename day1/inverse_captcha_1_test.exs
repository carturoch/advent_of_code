Code.load_file("inverse_captcha_1.exs", __DIR__)
ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule InverseCaptchaTest do
  use ExUnit.Case

  test "can process binaries or charlists" do
    assert InverseCaptcha.sum("123") == InverseCaptcha.sum(["1", "2", "3"], "1")
  end

  test "returns 0 for empty sequence" do
    assert InverseCaptcha.sum("1234") == 0
  end

  test "sums sequentially matching digits" do
    assert InverseCaptcha.sum("1122") == 3
  end

  test "sums sequentially matching digits repeatedly" do
    assert InverseCaptcha.sum("11115") == 3
  end

  test "sums the last element when matches the first one" do
    assert InverseCaptcha.sum("91212129") == 9
  end
end

