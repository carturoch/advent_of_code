Code.load_file("inverse_captcha_1.exs", __DIR__)
ExUnit.start
ExUnit.configure trace: true, exclude: :pending

defmodule InverseCaptchaTest do
  use ExUnit.Case

  describe "sum1" do
    test "returns 0 for empty sequence" do
      assert InverseCaptcha.sum1("1234") == 0
    end

    test "sums sequentially matching digits" do
      assert InverseCaptcha.sum1("1122") == 3
    end

    test "sums sequentially matching digits repeatedly" do
      assert InverseCaptcha.sum1("11115") == 3
    end

    test "sums the last element when matches the first one" do
      assert InverseCaptcha.sum1("91212129") == 9
    end
  end

  describe "sum2" do
    test "matches overflow" do
      assert InverseCaptcha.sum2("1212") == 6
      assert InverseCaptcha.sum2("1221") == 0
    end

    test "matches single pair" do
      assert InverseCaptcha.sum2("123425") == 4
    end
  end

  describe "is_match?" do
    def to_map(str) do
      (0..(String.length(str)))
      |> Enum.zip(String.graphemes(str))
      |> Enum.into(%{})
    end

    test "offset within str length" do
      assert InverseCaptcha.is_match?(to_map("ababbb"), 0, 2)
    end
    test "offset over the str length" do
      assert InverseCaptcha.is_match?(to_map("abaaba"), 4, 3)
      refute InverseCaptcha.is_match?(to_map("baaaba"), 4, 3)
      assert InverseCaptcha.is_match?(to_map("baaaab"), 5, 1)
    end
  end
end

